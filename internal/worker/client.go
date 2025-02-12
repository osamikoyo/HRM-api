package worker

import (
	"github.com/osamikoyo/hrm-api/internal/config"
	"github.com/osamikoyo/hrm-api/internal/worker/pb"
	"google.golang.org/grpc"
)

type Client struct{
	pb.WorkerServiceClient
}

func Init(cfg *config.Config) (*Client, error) {
	cc, err := grpc.Dial(cfg.WorkerSvcUrl, grpc.WithInsecure())
	if err != nil{
		return nil, err
	}

	return &Client{
		pb.NewWorkerServiceClient(cc),
	}, nil
}
