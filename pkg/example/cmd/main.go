package main

import (
	"context"
	"flag"
	"fmt"
	kitPrometheus "github.com/go-kit/kit/metrics/prometheus"
	_ "github.com/lib/pq"
	stdPrometheus "github.com/prometheus/client_golang/prometheus"
	"gt-kit/pkg/example"
	"gt-kit/shared/utils/config"
	"gt-kit/shared/utils/database"

	"github.com/go-kit/kit/log"

	"github.com/go-kit/kit/log/level"

	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func init()  {
	fmt.Println("Initiate Config")
	config.SetConfigFile("config", "pkg/example/config", "json")
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
		Subsystem: "example_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitPrometheus.NewSummaryFrom(stdPrometheus.SummaryOpts{
		Namespace: "api",
		Subsystem: "example_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitPrometheus.NewSummaryFrom(stdPrometheus.SummaryOpts{
		Namespace: "api",
		Subsystem: "example_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{}) // no fields here

	var srv example.Service
	{
		repository := example.NewRepo(database.NewDB(logger))
		srv = example.NewService(repository)
	}

	srv = example.LoggingMiddleware{Next: srv}
	srv = example.InstrumentingMiddleware{RequestCount: requestCount, RequestLatency: requestLatency, CountResult: countResult, Next: srv}

	endpoints := example.MakeEndpoints(srv)

	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := example.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}
