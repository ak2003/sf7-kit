package main

import (
	"context"
	"encoding/json"
	"gitlab.dataon.com/gophers/sf7-kit/pkg/custom_field"

	// "database/sql"
	"flag"
	"fmt"
	"net"

	"gitlab.dataon.com/gophers/sf7-kit/pkg/employee"
	"gitlab.dataon.com/gophers/sf7-kit/pkg/example"
	"gitlab.dataon.com/gophers/sf7-kit/pkg/example/model/protoc/model"
	"gitlab.dataon.com/gophers/sf7-kit/pkg/grpc_employee"
	modelEmp "gitlab.dataon.com/gophers/sf7-kit/pkg/grpc_employee/model/protoc/model"
	"gitlab.dataon.com/gophers/sf7-kit/pkg/leave"
	"gitlab.dataon.com/gophers/sf7-kit/pkg/user"
	"gitlab.dataon.com/gophers/sf7-kit/shared/connections"
	"gitlab.dataon.com/gophers/sf7-kit/shared/utils/config"
	"gitlab.dataon.com/gophers/sf7-kit/shared/utils/database"

	_ "github.com/denisenkom/go-mssqldb"
	kitPrometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	stdPrometheus "github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"

	"github.com/go-kit/kit/log"

	"github.com/go-kit/kit/log/level"

	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/cors"
)

func init() {
	fmt.Println("Initiate Config")
	config.SetConfigFile("config", "./config", "json")
}

func main() {
	// @todo port get from config
	var httpAddr = flag.String("http", ":8080", "http listen address")
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "user",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	var dbSlave *sqlx.DB
	{
		var err error
		dbSlave, err = connections.ConnSlave(logger)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}

	}
	var dbMaster *sqlx.DB
	{
		var err error
		dbMaster, err = connections.ConnMaster(logger)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}

	}

	flag.Parse()
	ctx := context.Background()

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger = log.NewLogfmtLogger(os.Stderr)
	fieldKeys := []string{"method", "error"}
	requestCount := kitPrometheus.NewCounterFrom(stdPrometheus.CounterOpts{
		Namespace: "api",
		Subsystem: "sf7_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitPrometheus.NewSummaryFrom(stdPrometheus.SummaryOpts{
		Namespace: "api",
		Subsystem: "sf7_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitPrometheus.NewSummaryFrom(stdPrometheus.SummaryOpts{
		Namespace: "api",
		Subsystem: "sf7_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{}) // no fields here

	// example package.
	var srv example.Service
	{
		repository := example.NewRepo(database.NewDB(logger))
		srv = example.NewService(repository)
	}

	srv = example.LoggingMiddleware{Next: srv}
	srv = example.InstrumentingMiddleware{RequestCount: requestCount, RequestLatency: requestLatency, CountResult: countResult, Next: srv}

	endpoints := example.MakeEndpoints(srv)

	// grpc employee pkg
	var srvGrpcEmployee grpc_employee.Service
	{
		repository := grpc_employee.NewRepo(dbSlave, dbMaster)
		srvGrpcEmployee = grpc_employee.NewService(repository)
	}

	srvGrpcEmployee = grpc_employee.LoggingMiddleware{Next: srvGrpcEmployee}
	srvGrpcEmployee = grpc_employee.InstrumentingMiddleware{RequestCount: requestCount, RequestLatency: requestLatency, CountResult: countResult, Next: srvGrpcEmployee}

	endpointsGrpcEmployee := grpc_employee.MakeEndpoints(srvGrpcEmployee)

	// user package
	var srvUser user.Service
	{
		repository := user.NewRepo(dbSlave, dbMaster, logger)
		srvUser = user.NewService(repository)
	}

	srvUser = user.LoggingMiddleware{Logger: logger, Next: srvUser}
	srvUser = user.InstrumentingMiddleware{RequestCount: requestCount, RequestLatency: requestLatency, CountResult: countResult, Next: srvUser}

	endpointsUser := user.MakeEndpoints(srvUser)

	var srvLeave leave.Service
	{
		repository := leave.NewRepo(dbSlave, dbMaster)

		srvLeave = leave.NewService(repository)
	}
	endpointsLeave := leave.MakeEndpoints(srvLeave)

	var srvEmployee employee.Service
	{
		repository := employee.NewRepo(dbSlave, dbMaster)

		srvEmployee = employee.NewService(repository)
	}
	endpointsEmployee := employee.MakeEndpoints(srvEmployee)


	var srvCustomField custom_field.Service
	{
		repository := custom_field.NewRepo(dbSlave, dbMaster)

		srvCustomField = custom_field.NewService(repository)
	}
	endpointsCustomField := custom_field.MakeEndpoints(srvCustomField)

	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := mux.NewRouter()
		// handler.Use(response.CommonMiddleware)
		handler.Use(commonMiddleware)
		handler = example.NewHTTPServer(ctx, endpoints, handler)
		handler = user.NewHTTPServer(ctx, endpointsUser, handler)
		handler = leave.NewHTTPServer(ctx, endpointsLeave, handler)
		handler = employee.NewHTTPServer(ctx, endpointsEmployee, handler)
		handler = grpc_employee.NewHTTPServer(ctx, endpointsGrpcEmployee, handler)
		handler = custom_field.NewHTTPServer(ctx, endpointsCustomField, handler)

		handlers := cors.Default().Handler(handler)
		// handler.Use(mux.CORSMethodMiddleware(handler))
		errs <- http.ListenAndServe(*httpAddr, handlers)
	}()

	// Starting RPC Server
	srvRpc := grpc.NewServer()
	model.RegisterExampleServer(srvRpc, srv)
	modelEmp.RegisterEmployeeServer(srvRpc, srvGrpcEmployee)

	go func() {
		level.Info(logger).Log("msg", "Starting RPC server at"+":7000")
		l, err := net.Listen("tcp", ":7000")
		if err != nil {
			level.Error(logger).Log("err", fmt.Errorf("could not listen to %s: %v", ":7000", err))
		}
		errs <- srvRpc.Serve(l)
	}()

	level.Error(logger).Log("exit", <-errs)
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

func HandshakeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		payload, _ := json.Marshal("OK")
		w.Write(payload)
	}
}
