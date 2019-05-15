package grpc

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"m15.io/kappa/pkg/delivery/grpc/conf_grpc"
	"m15.io/kappa/pkg/domain"
	"m15.io/kappa/pkg/usecases"
)

// server implements grpc service
type server struct {
	interactor usecases.ConfInteractor
}

// NewConfServerGrpc creates new server instance and registers it as grpc service
func NewConfServerGrpc(gserver *grpc.Server, interactor usecases.ConfInteractor) {
	confServer := &server{
		interactor: interactor,
	}

	conf_grpc.RegisterConfHandlerServer(gserver, confServer)
	reflection.Register(gserver)
}

// GetConf executes usecase GetConf method to fetch data from LDAP and create client app configuration for given username
func (s *server) GetConf(ctx context.Context, req *conf_grpc.FetchRequest) (*conf_grpc.Conf, error) {

	domainConf, err := s.interactor.GetConf(req.GetUsername())
	if err != nil {
		log.WithFields(log.Fields{
			"err":       err,
			"username":  req.GetUsername(),
			"ipaddr":    req.GetIpaddr(),
			"mac":       req.GetMac(),
			"timestamp": req.GetTimestamp(),
		}).Error("Error getting user conf")
		return nil, err
	}

	grpcConf := s.transformDomainGrpc(domainConf)

	log.WithFields(log.Fields{
		"username":  req.GetUsername(),
		"ipaddr":    req.GetIpaddr(),
		"mac":       req.GetMac(),
		"timestamp": req.GetTimestamp(),
		"conf":      grpcConf,
	}).Info("Succesfully created user conf")

	return grpcConf, nil
}

// transformDomainGrpc transforms domain.Conf to grpc.Conf
func (s *server) transformDomainGrpc(domainConf *domain.Conf) *conf_grpc.Conf {
	grpcConf := new(conf_grpc.Conf)
	grpcConf.Username = domainConf.Username

	grpcButtons := make([]*conf_grpc.Button, 0, 25)

	for _, domainButton := range domainConf.Buttons {
		b := &conf_grpc.Button{}
		b.Text = domainButton.Text
		b.Value = domainButton.Value
		grpcButtons = append(grpcButtons, b)
	}

	grpcConf.Button = grpcButtons

	return grpcConf
}
