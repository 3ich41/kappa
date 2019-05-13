package grpc

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	conf_grpc "m15.io/kappa/pkg/delivery/grpc/conf_grpc"
	"m15.io/kappa/pkg/usecases"
)

type server struct {
	usecase usecases.ConfUsecase
}

func NewConfServerGrpc(gserver *grpc.Server, confusecase usecases.ConfUsecase) {
	confServer := &server{
		usecase: confusecase,
	}

	conf_grpc.RegisterConfHandlerServer(gserver, confServer)
	reflection.Register(gserver)
}

func (s *server) GetConf(ctx context.Context, req *conf_grpc.FetchRequest) (*conf_grpc.Conf, error) {
	conf, err := s.usecase.GetConf(req.GetUsername())

}
