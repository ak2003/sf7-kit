package leave

//"github.com/form3tech-oss/jwt-go"
import (
	"context"

	"gitlab.dataon.com/gophers/sf7-kit/shared/response"

	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints, r *mux.Router) *mux.Router {

	//mw := jwtMiddleware.New(jwtMiddleware.Options{
	//	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
	//		return []byte(config.GetString("jwt.key")), nil
	//	},
	//	SigningMethod: jwt.SigningMethodHS256,
	//})

	// r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	apiv1 := api.PathPrefix("/v1").Subrouter()

	apiv1Sf7 := apiv1.PathPrefix("/sf7").Subrouter()
	apiv1Sf7Leave := apiv1Sf7.PathPrefix("/leave").Subrouter()
	// Add To cart
	apiv1Sf7Leave.Methods("POST").Path("/request").Handler(httpTransport.NewServer(
		endpoints.GetLeaveRequestListing,
		decodeGetLeaveRequestListingReq,
		response.EncodeJsonWithStatusCode,
	))

	apiv1Sf7Leave.Methods("POST").Path("/filterRequest").Handler(httpTransport.NewServer(
		endpoints.GetLeaveRequestFilterListing,
		decodeGetLeaveRequestFilterListingReq,
		response.EncodeJsonWithStatusCode,
	))

	r.Methods("GET").Path("/metrics").Handler(promhttp.Handler())

	// r.Handle("/api/{rest:.*}", HandshakeHandler()).Methods("OPTIONS")

	return r

}

// func HandshakeHandler() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set("Access-Control-Allow-Headers", "*")

// 		payload, _ := json.Marshal("OK")
// 		w.Write(payload)
// 	}
// }
