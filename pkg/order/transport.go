package order

//"github.com/form3tech-oss/jwt-go"
import (
	"context"
	"gitlab.dataon.com/gophers/sf7-kit/shared/response"
	"net/http"

	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {

	//mw := jwtMiddleware.New(jwtMiddleware.Options{
	//	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
	//		return []byte(config.GetString("jwt.key")), nil
	//	},
	//	SigningMethod: jwt.SigningMethodHS256,
	//})

	r := mux.NewRouter()
	r.Use(commonMiddleware)

	v1 := r.PathPrefix("/v1").Subrouter()


	// Add To cart
	v1.Methods("POST").Path("/order/cart").Handler(httpTransport.NewServer(
		endpoints.AddToCart,
		decodeAddToCartReq,
		response.EncodeJson,
	))

	// Cart List

	// Delete cart by product_id
	v1.Methods("DELETE").Path("/order/cart").Handler(httpTransport.NewServer(
		endpoints.DeleteItemCart,
		decodeDelItemCartReq,
		response.EncodeJson,
	))

	// Empty Cart (Delete by cartID)


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
