package vocation

import (
	"github.com/osamikoyo/hrm-api/internal/config"
	"github.com/osamikoyo/hrm-api/internal/vocation/pb"
	"google.golang.org/grpc"
)

type Client struct{
	client pb.VocationServiceClient
}

func InitClient(cfg *config.Config) (*Client, error) {
	cc, err := grpc.Dial(cfg.VocationSrcUrl, grpc.WithInsecure())
	if err != nil{
		return nil, err
	}

	return &Client{
		pb.NewVocationServiceClient(cc),
	}, nil
}
