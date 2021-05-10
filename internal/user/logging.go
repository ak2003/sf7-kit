package user

import (
	"context"
	"github.com/go-kit/kit/log"
	"time"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next   Service
}

func (mw LoggingMiddleware) CreateUser(ctx context.Context, email string, password string) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "CreateUser",
			"input", email,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.CreateUser(ctx, email, password)
	return
}

func (mw LoggingMiddleware) LoginUser(ctx context.Context, email string, password string) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "LoginUser",
			"input", email,
			"output", "token",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.LoginUser(ctx, email, password)
	return
}


func (mw LoggingMiddleware) GetUser(ctx context.Context, id string, token string) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "GetUser",
			"input", id,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.GetUser(ctx, id, token)
	return
}