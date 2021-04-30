#!/bin/bash

# /bin/bash ./scripts/create-endpoint.sh $1 $2 $3
if [ $# -eq 4 ]
then
    echo "creating API..."
else
    echo "invalid argument please pass two arguments "
    exit
fi

SERVICENAME=$1
FUNC=$2
METHOD=$3
APIPATH=$4

SERVICE="pkg/$SERVICENAME"
FUNC="$(tr '[:lower:]' '[:upper:]' <<< ${FUNC:0:1})${FUNC:1}"
REQUEST=$FUNC"Request"
RESPONSE=$FUNC"Response"

# Validation
# - service name must be exist
if ! [ -d "./$SERVICE" ]
then
    echo "service $SERVICE does not exists."
    exit 0
fi
# - function name must be not available
if grep -q "$FUNC" "$SERVICE/service.go"; then
  echo "Function name already exist"
  exit 0
fi
# - method with option [GET,POST,DELETE,PUT]
if ! [[ "$METHOD" = "POST" || "$METHOD" = "GET" || "$METHOD" = "PUT" || "$METHOD" = "DELETE" ]]
then
  echo "Method Not Allow"
  exit 0
fi
# - apipath must be not used
if grep -q "$APIPATH" "$SERVICE/transport.go"; then
  echo "API Path already exist"
  exit 0
fi

#generate protoc
# install protobuf for mac : brew install protobuf -> https://grpc.io/docs/languages/go/quickstart/
PATH=$PATH:$GOPATH/bin/ protoc --go_out=./$SERVICE/model ./$SERVICE/model/$FUNC.proto
#
INTERFACE="$FUNC(ctx context.Context, req *model.$REQUEST) (interface{}, error)"

#create interface service
VAR=$(wc -l < $SERVICE/service.go)
#echo "LINE $VAR"
#
sed -i "" ''"$VAR"' i\
'"$INTERFACE"'' $SERVICE/service.go
gofmt -w $SERVICE/service.go

#create endpoint struct
ENDPOINT="$FUNC endpoint.Endpoint"
lineNumber=$(awk '/Endpoints struct/{ print NR; exit }' $SERVICE/endpoint.go)
awk "NR >= $lineNumber" $SERVICE/endpoint.go |
while read -r line; do
  if [ "$line" = "}" ]
  then
    sed -i "" ''"$lineNumber"' i\
    '"$ENDPOINT"'' $SERVICE/endpoint.go
    gofmt -w $SERVICE/endpoint.go
    break
  fi
  lineNumber=$((lineNumber+1))
done

#create endpoint MakeEndpoints
ENDPOINT=''"$FUNC"': make'"$FUNC"'Endpoint(s),\
 '
lineNumber=$(awk '/MakeEndpoints/{ print NR; exit }' $SERVICE/endpoint.go)

awk "NR >= $lineNumber" $SERVICE/endpoint.go |
while read -r line; do
  if [ "$line" = "}" ]
  then
    sed -i "" ''"$lineNumber"' i\
    '"$ENDPOINT"'' $SERVICE/endpoint.go
    gofmt -w $SERVICE/endpoint.go
    break
  fi
  lineNumber=$((lineNumber+1))
done

## endpoint.go
gofmt -w $SERVICE/endpoint.go
# shellcheck disable=SC2129
echo " " >> $SERVICE/endpoint.go
echo "func make"$FUNC"Endpoint(s Service) endpoint.Endpoint {" >> $SERVICE/endpoint.go
echo "return func(ctx context.Context, request interface{}) (interface{}, error) {" >> $SERVICE/endpoint.go
echo "req := request.(*model."$REQUEST")" >> $SERVICE/endpoint.go
echo "resp, err := s."$FUNC"(ctx, req)" >> $SERVICE/endpoint.go
echo "responseBody := response.Body{Data: resp}" >> $SERVICE/endpoint.go
echo "return response.CreateResponse{RespBody: responseBody, Err: err}, nil" >> $SERVICE/endpoint.go
echo "}" >> $SERVICE/endpoint.go
echo "}" >> $SERVICE/endpoint.go

gofmt -w $SERVICE/endpoint.go

## logic.go
gofmt -w $SERVICE/logic.go
# shellcheck disable=SC2129
echo " " >> $SERVICE/logic.go
echo "func (s service) "$FUNC"(ctx context.Context, req *model."$FUNC"Request) (interface{}, error) {" >> $SERVICE/logic.go
echo "return nil, nil" >> $SERVICE/logic.go
echo "}" >> $SERVICE/logic.go

gofmt -w $SERVICE/logic.go

# Instrumenting.go
gofmt -w $SERVICE/instrumenting.go
# shellcheck disable=SC2129
echo " " >> $SERVICE/instrumenting.go
echo "func (mw InstrumentingMiddleware) "$FUNC"(ctx context.Context, req *model."$FUNC"Request) (output interface{}, err error) {" >> $SERVICE/instrumenting.go
echo "defer func(begin time.Time) {" >> $SERVICE/instrumenting.go
echo 'lvs := []string{"method", "'$FUNC'", "error", fmt.Sprint(err != nil)}' >> $SERVICE/instrumenting.go
echo "mw.RequestCount.With(lvs...).Add(1)" >> $SERVICE/instrumenting.go
echo "mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())" >> $SERVICE/instrumenting.go
echo "}(time.Now())" >> $SERVICE/instrumenting.go
echo "" >> $SERVICE/instrumenting.go
echo "output, err = mw.Next."$FUNC"(ctx, req)" >> $SERVICE/instrumenting.go
echo "return" >> $SERVICE/instrumenting.go
echo "}" >> $SERVICE/instrumenting.go

gofmt -w $SERVICE/instrumenting.go

# logging.go
gofmt -w $SERVICE/logging.go
# shellcheck disable=SC2129
echo " " >> $SERVICE/logging.go
echo "func (mw LoggingMiddleware) "$FUNC"(ctx context.Context, req *model."$FUNC"Request) (output interface{}, err error) {" >> $SERVICE/logging.go
echo 'msg := "Incoming Request"' >> $SERVICE/logging.go
echo "defer func(begin time.Time) {" >> $SERVICE/logging.go
echo "fields := make(map[string]interface{})" >> $SERVICE/logging.go
echo 'fields["input"] = req' >> $SERVICE/logging.go
echo 'fields["output"] = output' >> $SERVICE/logging.go
echo 'fields["err"] = err' >> $SERVICE/logging.go
echo "if err != nil {" >> $SERVICE/logging.go
echo "logger.Error(fields, msg)" >> $SERVICE/logging.go
echo "return" >> $SERVICE/logging.go
echo "" >> $SERVICE/logging.go
echo "}" >> $SERVICE/logging.go
echo "logger.Info(fields, msg)" >> $SERVICE/logging.go
echo "}(time.Now())" >> $SERVICE/logging.go
echo "" >> $SERVICE/logging.go
echo "return mw.Next."$FUNC"(ctx, req)" >> $SERVICE/logging.go
echo "}" >> $SERVICE/logging.go

gofmt -w $SERVICE/logging.go

# reqresp.go
gofmt -w $SERVICE/reqresp.go
# shellcheck disable=SC2129
echo " " >> $SERVICE/reqresp.go
echo "func decode"$FUNC"Req(_ context.Context, r *http.Request) (interface{}, error) {" >> $SERVICE/reqresp.go
echo "var req *model."$FUNC"Request" >> $SERVICE/reqresp.go
echo "err := json.NewDecoder(r.Body).Decode(&req)" >> $SERVICE/reqresp.go
echo "if err != nil {" >> $SERVICE/reqresp.go
echo "return nil, err" >> $SERVICE/reqresp.go
echo "}" >> $SERVICE/reqresp.go
echo "return req, nil" >> $SERVICE/reqresp.go
echo "}" >> $SERVICE/reqresp.go
#
gofmt -w $SERVICE/reqresp.go

# Transport.go
lineNumber=$(awk '/addNew endpoint/{ print NR; exit }' $SERVICE/transport.go)
echo "$lineNumber"

lineNumber=$((lineNumber - 1))
script='v1.Methods("'"$METHOD"'").Path("/'"$APIPATH"'").Handler(httpTransport.NewServer(\
endpoints.'"$FUNC"',\
decode'"$FUNC"'Req,\
response.EncodeResponse,\
opts...,\
))\
 '
sed -i "" ''"$lineNumber"' i\
    '"$script"'' $SERVICE/transport.go
gofmt -w $SERVICE/transport.go

echo "API generated."