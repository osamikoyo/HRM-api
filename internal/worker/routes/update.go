package routes

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/osamikoyo/hrm-api/internal/worker/pb"
)

func Update(ctx echo.Context, c pb.WorkerServiceClient) error {
	var req UpdateReq

	if err := ctx.Bind(&req);err != nil{
		return err
	}

	resp, err := c.Update(context.Background(), &pb.UpdateWorkerRequest{
		UserID: req.ID,
		NewWorkerParametres: ToPB(&req.Worker),
	})

	if err != nil{
		return fmt.Errorf("error to delete req: %w", err)
	}

	if resp.Status != http.StatusOK{
		return fmt.Errorf("internal grpc server error: %w", errors.New(resp.Error))
	}

	return nil
}