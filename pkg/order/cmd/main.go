package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	kitPrometheus "github.com/go-kit/kit/metrics/prometheus"
	_ "github.com/lib/pq"
	stdPrometheus "github.com/prometheus/client_golang/prometheus"
	"gt-kit/pkg/order"
	"gt-kit/shared/utils/config"

	"github.com/go-kit/kit/log"

	"github.com/go-kit/kit/log/level"

	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var serviceName = "order"

func init()  {
	fmt.Println("Initiate Config")
	config.SetConfigFile("config", serviceName + "/config", "json")
}

func main() {

	var httpAddr = flag.String("http", ":7070", "http listen address")
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
			dbUser = config.GetDBUser(dbDriver)
			dbPass = config.GetDBPass(dbDriver)
			dbHost = config.GetDBHost(dbDriver)
			dbPort = config.GetDBPort(dbDriver)
			dbName = config.GetDBName(dbDriver)
		)
		var dbSource = "postgresql://"+ dbUser +":"+ dbPass +"@"+dbHost+":"+dbPort+"/"+ dbName+"?sslmode=disable"
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

	var srv order.Service
	{
		repository := order.NewRepo(db)
		srv = order.NewService(repository)
	}

	srv = order.LoggingMiddleware{Logger: logger, Next: srv}
	srv = order.InstrumentingMiddleware{RequestCount: requestCount, RequestLatency: requestLatency, CountResult: countResult, Next: srv}

	endpoints := order.MakeEndpoints(srv)

	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := order.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}
