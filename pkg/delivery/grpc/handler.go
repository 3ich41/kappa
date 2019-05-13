package grpc

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	conf_grpc "m15.io/kappa/pkg/delivery/grpc/conf_grpc"
	"m15.io/kappa/pkg/domain"
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
	grpcConf := new(conf_grpc.Conf)

	domainConf, err := s.usecase.GetConf(req.GetUsername())
	if err != nil {
		return grpcConf, err
	}

	grpcConf.Username = domainConf.Username

	return grpcConf, nil
}

func (s *server) transformDomainGrpc(domainConf *domain.Conf) *conf_grpc.Conf {
	grpcConf := new(conf_grpc.Conf)
	grpcConf.Username = domainConf.Username

	grpcButtons := make([]conf_grpc.Button, 0, 25)

	for domainButton := range domainConf.Buttons {

	}

	return grpcConf
}
