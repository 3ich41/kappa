package main

import (
	"fmt"
	"net"
	"os"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"m15.io/kappa/pkg/config"
	confgrpc "m15.io/kappa/pkg/delivery/grpc"
	"m15.io/kappa/pkg/infrastructure"
	"m15.io/kappa/pkg/repositories"
	"m15.io/kappa/pkg/usecases"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

func main() {
	c := config.NewConfig()

	switch c.LogLevel {
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "TRACE":
		log.SetLevel(log.TraceLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}

	ldapclient := &infrastructure.LDAPClient{
		Base:         c.LdapBase,
		Host:         c.LdapHost,
		Port:         c.LdapPort,
		UseSSL:       c.LdapUseSSL,
		SkipTLS:      c.LdapSkipTLS,
		BindDN:       c.LdapBindDN,
		BindPassword: c.LdapBindPassword,
		UserFilter:   c.LdapUserFilter,
	}

	ldapGroupRepository := repositories.NewLdapGroupRepository(ldapclient)

	confInteractor := usecases.NewConfInteractor(ldapGroupRepository)

	address := fmt.Sprintf(":%d", c.GrpcPort)
	l, err := net.Listen("tcp", address)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	server := grpc.NewServer()
	confgrpc.NewConfServerGrpc(server, confInteractor)
	log.Infof("Server Run at :%d", c.GrpcPort)

	err = server.Serve(l)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}
