package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"net"

	"github.com/go-kit/kit/log"
	kitPrometheus "github.com/go-kit/kit/metrics/prometheus"
	_ "github.com/lib/pq"
	stdPrometheus "github.com/prometheus/client_golang/prometheus"
	"gitlab.dataon.com/gophers/sf7-kit/pkg/product"
	"gitlab.dataon.com/gophers/sf7-kit/pkg/product/model/protoc/model"
	"gitlab.dataon.com/gophers/sf7-kit/shared/utils/config"
	"google.golang.org/grpc"

	"github.com/go-kit/kit/log/level"

	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var serviceName = "product"

func init() {
	fmt.Println("Initiate Config")
	config.SetConfigFile("config", "pkg/"+serviceName+"/config", "json")
}

func main() {

	var httpAddr = flag.String("http", ":8080", "http listen address")
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", serviceName,
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	var db *sql.DB
	{
		var err error
		var (
			dbDriver = "postgresql"
			dbUser   = config.GetDBUser(dbDriver)
			dbPass   = config.GetDBPass(dbDriver)
			dbHost   = config.GetDBHost(dbDriver)
			dbPort   = config.GetDBPort(dbDriver)
			dbName   = config.GetDBName(dbDriver)
		)
		var dbSource = "postgresql://" + dbUser + ":" + dbPass + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=disable"
		level.Info(logger).Log("dbInfo", dbSource)
		db, err = sql.Open("postgres", dbSource)
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
		Subsystem: serviceName + "_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitPrometheus.NewSummaryFrom(stdPrometheus.SummaryOpts{
		Namespace: "api",
		Subsystem: serviceName + "_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitPrometheus.NewSummaryFrom(stdPrometheus.SummaryOpts{
		Namespace: "api",
		Subsystem: serviceName + "_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{}) // no fields here

	var srv product.Service
	{
		repository := product.NewRepo(db)
		srv = product.NewService(repository)
	}

	srv = product.LoggingMiddleware{Logger: logger, Next: srv}
	srv = product.InstrumentingMiddleware{RequestCount: requestCount, RequestLatency: requestLatency, CountResult: countResult, Next: srv}

	endpoints := product.MakeEndpoints(srv)

	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := product.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	// Starting RPC Server
	srvRpc := grpc.NewServer()
	model.RegisterProductsServer(srvRpc, srv)

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
