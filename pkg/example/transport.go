package example

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gitlab.dataon.com/gophers/sf7-kit/shared/response"
)

func NewHTTPServer(_ context.Context, endpoints Endpoints, r *mux.Router) *mux.Router {
	var logger log.Logger
	opts := []httpTransport.ServerOption{
		httpTransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httpTransport.ServerErrorEncoder(response.EncodeError),
	}

	v1 := r.PathPrefix("/v1").Subrouter()

	// HealthCheck endpoint
	v1.Methods("POST").Path("/health-check").Handler(httpTransport.NewServer(
		endpoints.HealthCheck,
		decodeHealthCheckReq,
		response.EncodeResponse,
		opts...,
	))

	/*addNew endpoint -> @todo: don't remove*/

	// Metric endpoint
	r.Methods("GET").Path("/metrics").Handler(promhttp.Handler())
	return r
}
