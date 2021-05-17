package user

//"github.com/form3tech-oss/jwt-go"
import (
	"context"
	"net/http"

	"github.com/form3tech-oss/jwt-go"
	"gitlab.dataon.com/gophers/sf7-kit/shared/response"
	"gitlab.dataon.com/gophers/sf7-kit/shared/utils/config"

	jwtMiddleware "github.com/auth0/go-jwt-middleware"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/negroni"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints, r *mux.Router) *mux.Router {

	mw := jwtMiddleware.New(jwtMiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetString("jwt.key")), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	// r.Use(commonMiddleware)

	v1 := r.PathPrefix("/v1").Subrouter()

	// User Registration
	v1.Methods("POST").Path("/user").Handler(httpTransport.NewServer(
		endpoints.CreateUser,
		decodeUserReq,
		response.EncodeResponse,
	))

	// User Login
	v1.Methods("POST").Path("/user/login").Handler(httpTransport.NewServer(
		endpoints.LoginUser,
		decodeLoginReq,
		response.EncodeResponse,
	))

	//Auth Require
	ar := mux.NewRouter()
	arV1 := ar.PathPrefix("/v1").Subrouter()
	arV1.Methods("GET").Path("/profile/user/{id}").Handler(httpTransport.NewServer(
		endpoints.GetUser,
		decodeEmailReq,
		response.EncodeResponse,
	))
	an := negroni.New(negroni.HandlerFunc(mw.HandlerWithNext), negroni.Wrap(ar))
	v1.PathPrefix("/profile").Handler(an)

	// Metric
	r.Methods("GET").Path("/metrics").Handler(promhttp.Handler())

	return r

}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
