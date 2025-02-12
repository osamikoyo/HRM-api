package worker

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/osamikoyo/hrm-api/internal/config"
	"github.com/osamikoyo/hrm-api/internal/worker/routes"
)

func RegisterRoutes(e *echo.Echo, cfg *config.Config) error {
	client, err := InitClient(cfg)
	if err != nil{
		return err
	}

	workers := e.Group("/worker")
	workers.POST("/add", client.Add)
	
	return nil
}

func (c *Client) Add(ctx echo.Context) error {
	return routes.Add(ctx, c.client)
}

func (c *Client) Get(ctx echo.Context) error {
	result, err := routes.Get(ctx, c.client)
	if err != nil{
		return err
	}

	return ctx.JSON(http.StatusOK, result)
}

func (c *Client) Update(ctx echo.Context) error {
	return routes.Update(ctx, c.client)
}

func (c *Client) Delete(ctx echo.Context) error {
	return routes.Delete(ctx, c.client)
}