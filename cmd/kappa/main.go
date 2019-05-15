package main

import (
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"
	"m15.io/kappa/pkg/config"
	confgrpc "m15.io/kappa/pkg/delivery/grpc"
	"m15.io/kappa/pkg/infrastructure"
	"m15.io/kappa/pkg/repositories"
	"m15.io/kappa/pkg/usecases"

	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
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
		Base:         "dc=resspu,dc=t00,dc=mil,dc=pl",
		Host:         "32.1.4.2",
		Port:         389,
		UseSSL:       false,
		SkipTLS:      true,
		BindDN:       "cn=test,ou=T00_USERS,dc=resspu,dc=t00,dc=mil,dc=pl",
		BindPassword: "1234qwerASDF",
		UserFilter:   "(sAMAccountName=%s)",
		GroupFilter:  "(sAMAccountName=%s)",
		Attributes:   []string{},
	}

	ldapGroupRepository := repositories.NewLdapGroupRepository(ldapclient)

	confInteractor := usecases.NewConfInteractor(ldapGroupRepository)

	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	server := grpc.NewServer()
	confgrpc.NewConfServerGrpc(server, confInteractor)
	fmt.Println("Server Run at :50051")

	err = server.Serve(l)
	if err != nil {
		fmt.Println("Unexpected Error", err)
	}
}
