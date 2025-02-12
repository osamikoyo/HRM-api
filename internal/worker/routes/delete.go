package routes

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/osamikoyo/hrm-api/internal/worker/pb"
)


func Delete(ctx echo.Context, c pb.WorkerServiceClient) error {
	var DeleteRequest DeleteRequest
	if err := ctx.Bind(&DeleteRequest);err != nil{
		return err
	}

	resp, err := c.Delete(context.Background(), &pb.DeleteWorkerRequest{
		UserID: DeleteRequest.ID,
	})

	if err != nil{
		return fmt.Errorf("error to delete req: %w", err)
	}

	if resp.Status != http.StatusOK{
		return fmt.Errorf("internal grpc server error: %w", errors.New(resp.Error))
	}

	return nil
}