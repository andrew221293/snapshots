package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/andrew221293/snapshots/internal/entity"
	"github.com/andrew221293/snapshots/internal/transport"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var genericErrResponse = entity.ResponseError{
	Error: "something went wrong",
	Code:  "959a1908-62f0-4cad-afc0-d9b4300085db",
}

func main() {
	ctx := context.Background()
	address := os.Getenv("BACKEND_HOST")

	e := echo.New()
	e.HTTPErrorHandler = customErrorHandler
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())

	router := &transport.Router{
		Echo:    e,
		Address: address,
		Handler: transport.EchoHandler{},
	}

	go func() {
		if err := router.Start(); err != nil {
			e.Logger.Infof("Shutting down ser")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func customErrorHandler(err error, e echo.Context) {
	var ce entity.CustomError
	if errors.As(err, &ce) {
		e.JSON(ce.HTTPCode, ce.ToResponseError()) // nolint: errcheck
		return
	}
	e.JSON(http.StatusInternalServerError, genericErrResponse) // nolint: errcheck
}
