package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/osamikoyo/hrm-api/internal/vocation/pb"
)

func Add(c echo.Context, cc pb.VocationServiceClient) error {
	var voc Vocation
	if err := c.Bind(&voc);err != nil{
		return err
	}

	resp, err := cc.Add(context.Background(), &pb.AddVocationRequest{
		Vocation: ToPB(&voc),
	})
	if err != nil{
		return fmt.Errorf("cant do a request: %w", err)
	}

	if resp.Status != http.StatusOK{
		return fmt.Errorf("internal service error: %s", resp.Error)
	}

	return nil
}