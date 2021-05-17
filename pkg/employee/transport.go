package employee

//"github.com/form3tech-oss/jwt-go"
import (
	"context"
	"net/http"

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
	r.Use(commonMiddleware)
	api := r.PathPrefix("/api").Subrouter()
	apiv1 := api.PathPrefix("/v1").Subrouter()

	apiv1Sf7 := apiv1.PathPrefix("/sf7").Subrouter()
	apiv1Sf7Employee := apiv1Sf7.PathPrefix("/employee").Subrouter()
	// Add To cart
	apiv1Sf7Employee.Methods("POST").Path("/information").Handler(httpTransport.NewServer(
		endpoints.GetEmployeeInformation,
		decodeGetEmployeeInformationReq,
		response.EncodeJsonWithStatusCode,
	))

	apiv1Sf7Employee.Methods("POST").Path("/editInformation").Handler(httpTransport.NewServer(
		endpoints.GetEmployeeEditInformation,
		decodeGetEmployeeEditInformationReq,
		response.EncodeJsonWithStatusCode,
	))

	apiv1Sf7Employee.Methods("POST").Path("/masterAddress").Handler(httpTransport.NewServer(
		endpoints.GetEmployeeMasterAddress,
		decodeGetEmployeeMasterAddressReq,
		response.EncodeJsonWithStatusCode,
	))

	apiv1Sf7Employee.Methods("POST").Path("/updateAddress").Handler(httpTransport.NewServer(
		endpoints.UpdateEmployeeMasterAddress,
		decodeGetEmployeeUpdateAddressReq,
		response.EncodeJsonWithStatusCode,
	))

	apiv1Sf7Employee.Methods("POST").Path("/createAddress").Handler(httpTransport.NewServer(
		endpoints.CreateEmployeeMasterAddress,
		decodeGetEmployeeCreateAddressReq,
		response.EncodeJsonWithStatusCode,
	))

	apiv1Sf7Employee.Methods("POST").Path("/getCity").Handler(httpTransport.NewServer(
		endpoints.GetCity,
		decodeGetCityReq,
		response.EncodeJsonWithStatusCode,
	))

	apiv1Sf7Employee.Methods("POST").Path("/getAddressType").Handler(httpTransport.NewServer(
		endpoints.GetAddressType,
		decodeGetAddressTypeReq,
		response.EncodeJsonWithStatusCode,
	))

	apiv1Sf7Employee.Methods("POST").Path("/getOwnerStatus").Handler(httpTransport.NewServer(
		endpoints.GetOwnerStatus,
		decodeGetOwnerStatusReq,
		response.EncodeJsonWithStatusCode,
	))

	apiv1Sf7Employee.Methods("POST").Path("/getStayStatus").Handler(httpTransport.NewServer(
		endpoints.GetStayStatus,
		decodeGetStayStatusReq,
		response.EncodeJsonWithStatusCode,
	))

	r.Methods("GET").Path("/metrics").Handler(promhttp.Handler())

	// r.Handle("/api/{rest:.*}", HandshakeHandler()).Methods("OPTIONS")

	return r

}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Token, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		next.ServeHTTP(w, r)
	})
}

// func HandshakeHandler() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set("Access-Control-Allow-Headers", "*")

// 		payload, _ := json.Marshal("OK")
// 		w.Write(payload)
// 	}
// }
