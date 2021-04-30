#!/bin/bash

#marker="// Metric endpoint"

#lineNumber=$(awk '/addNew endpoint/{ print NR; exit }' example/transport.go)
#echo "$lineNumber"

#lineNumber=$((lineNumber - 1))
#script='v1.Methods("POST").Path("/health-check").Handler(httpTransport.NewServer(\
#endpoints.HealthCheck,\
#decodeHealthCheckReq,\
#response.EncodeResponse,\
#opts...,\
#))\
# '
#sed -i "" ''"$lineNumber"' i\
#    '"$script"'' example/transport.go

#sed '35 i New Line with sed' example/transport.go

SERVICE="example"
FUNC="User"
REQUEST="AddToCartRequest"
RESPONSE="AddToCartResponse"
PATHNAME="user/get4"
#
INTERFACE="$FUNC(ctx context.Context, param *model.$REQUEST) (*model.$RESPONSE, error)"
#
#VAR=$(wc -l < $SERVICE/service.go)
#echo "LINE $VAR"
#
# shellcheck disable=SC1004
#sed -i "" '/^anothervalue=two/i'' i\
#'"$INTERFACE"'' $SERVICE/service.go

#sed -i'.ex' -e '/^anothervalue=.*/i before=me' $SERVICE/transport.go
#sed -i "" 's/two/&new\
#/' $SERVICE/transport.go


#if grep -q "user/get3" "example/transport.go"; then
#  echo "Find"
#  exit 0
#fi
#echo "empty"

#if ! [ -d "./$SERVICE" ]
#then
#    echo "service $SERVICE does not exists."
#    exit 0
#fi
#METHOD="PUT"
#if ! [[ "$METHOD" = "POST" || "$METHOD" = "GET" ]]
#then
#  echo "A"
#  exit 0
#fi
#if grep -q "$PATHNAME" "$SERVICE/transport.go"; then
#  echo "Path already exist"
#  exit 0
#fi

foo=$1
foo="$(tr '[:lower:]' '[:upper:]' <<< ${foo:0:1})${foo:1}"
echo $foo;
