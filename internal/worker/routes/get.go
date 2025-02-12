package routes

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/osamikoyo/hrm-api/internal/worker/pb"
)


func Get(ctx echo.Context, c pb.WorkerServiceClient) (*Worker, error) {
	var req GetRequest
	if err := ctx.Bind(&req);err != nil{
		return nil, err
	}

	resp, err := c.Get(context.Background(), &pb.GetWorkerRequest{
		UserID: req.ID,
	})

	if err != nil{
		return nil, fmt.Errorf("error to delete req: %w", err)
	}

	if resp.Response.Status != http.StatusOK{
		return nil, fmt.Errorf("internal grpc server error: %w", errors.New(resp.Response.Error))
	}

	return ToModels(resp.Worker), nil
}