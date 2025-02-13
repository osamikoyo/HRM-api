package routes

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/osamikoyo/hrm-api/internal/vocation/pb"
)

func Get(c echo.Context, cc pb.VocationServiceClient) (*Vocation, error) {
	var req GetReq

	if err := c.Bind(&req);err != nil{
		return nil, err
	}

	resp, err := cc.Get(context.Background(), &pb.GetVocationRequest{
		UserID: req.UserID,
	})

	if err != nil{
		return nil, fmt.Errorf("cant do a request: %w", err)
	}

	voc, err := ToModels(resp.Vocation)
	if err != nil{
		return nil, fmt.Errorf("error to bind to models: %w", err)
	}

	return voc, nil
}