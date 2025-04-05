package server

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type API struct {
	addr string
	db   *sql.DB
}

func NewAPI(addr string, db *sql.DB) *API {
	return &API{
		addr: addr,
		db:   db,
	}
}

func (api *API) Run() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong!")
	})

	e.Logger.Fatal(e.Start(api.addr))
}
