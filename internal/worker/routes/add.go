package routes

import (
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/osamikoyo/hrm-api/internal/worker/pb"
)

func Add(ctx echo.Context, c pb.WorkerServiceClient) error {
	var worker Worker
	if err := ctx.Bind(&worker);err != nil{
		return err
	}

	resp, err := c.Add(context.Background(), &pb.AddWorkerRequest{
		Worker: ToPB(&worker),
	})
	if err != nil{
		
	}

	if resp.Respone.Status != http.StatusOK{
		return errors.New("internal server error")
	}

	return nil
}
