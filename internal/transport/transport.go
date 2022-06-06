package transport

import (
	"fmt"
	"github.com/andrew221293/snapshots/internal/entity"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
)

type (
	Router struct {
		*echo.Echo
		Address string
		Handler EchoHandler
	}
	EchoHandler struct{}
)

//LocalHost routing
func (r *Router) Start() error {
	base := r.Group("/custom-endpoints")
	user := os.Getenv("BASIC_AUTH_USER")
	pass := os.Getenv("BASIC_AUTH_PASS")

	base.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == user && password == pass {
			return true, nil
		}
		return false, entity.CustomError{
			Err:      fmt.Errorf("basic auth failed"),
			HTTPCode: http.StatusUnauthorized,
			Code:     "e6807c42-3568-41de-a15f-fe0f073ab657",
		}
	}))

	//snapshots endpoints
	snapshots := base.Group("/snapshots")
	snapshots.GET("", r.Handler.Snapshots)

	return r.Echo.Start(r.Address)
}
