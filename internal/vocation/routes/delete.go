package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/osamikoyo/hrm-api/internal/vocation/pb"
)

func Delete(c echo.Context, cc pb.VocationServiceClient) error {
	var req DeleteReq
	if err := c.Bind(&req);err != nil{
		return err
	}

	resp, err := cc.Delete(context.Background(), &pb.DeleteVocationRequest{
		UserID: req.UserID,
	})

	if err != nil{
		return fmt.Errorf("cant do a request: %w", err)
	}

	if resp.Status != http.StatusOK{
		return fmt.Errorf("internal service error: %s", resp.Error)
	}

	return nil
}