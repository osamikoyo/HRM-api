package vocation

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/osamikoyo/hrm-api/internal/config"
	"github.com/osamikoyo/hrm-api/internal/vocation/routes"
)

func RegisterRoutes(e *echo.Echo, cfg  *config.Config) {
	cc, err := InitClient(cfg)
	if err != nil{
		e.Logger.Panic(err)
	}

	vocs := e.Group("/vocation")
	vocs.Use(middleware.Logger())

	vocs.POST("/add", cc.Add)
	vocs.POST("/get", cc.Get)
	vocs.POST("/delete", cc.Delete)
}

func (c *Client) Add(ctx echo.Context) error {
	err := routes.Add(ctx, c.client)
	if err != nil{
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.String(http.StatusCreated, "success")
}

func (c *Client) Get(ctx echo.Context) error {
	resp, err := routes.Get(ctx, c.client)
	if err != nil{
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (c *Client) Delete(ctx echo.Context) error {
	err := routes.Delete(ctx, c.client)
	if err != nil{
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.String(http.StatusOK, "success")
}