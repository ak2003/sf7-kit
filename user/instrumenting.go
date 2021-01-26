package user

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

type InstrumentingMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	CountResult    metrics.Histogram
	Next           Service
}

func (mw InstrumentingMiddleware) CreateUser(ctx context.Context, email string, password string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "CreateUser", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.CreateUser(ctx, email, password)
	return
}

func (mw InstrumentingMiddleware) LoginUser(ctx context.Context, email string, password string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "LoginUser", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.Next.LoginUser(ctx, email, password)
	return
}

func (mw InstrumentingMiddleware) GetUser(ctx context.Context, id string, token string) (n string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "count", "error", "false"}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
		//mw.countResult.Observe(float64(n))
	}(time.Now())

	n, err = mw.Next.GetUser(ctx, id, token)
	return
}
