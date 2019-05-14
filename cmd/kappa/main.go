package main

import (
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"
	confgrpc "m15.io/kappa/pkg/delivery/grpc"
	"m15.io/kappa/pkg/infrastructure"
	"m15.io/kappa/pkg/repositories"
	"m15.io/kappa/pkg/usecases"

	log "github.com/sirupsen/logrus"
)

func main() {
	//c := config.NewConfig()

	//fmt.Println(c)

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

	usecase := usecases.NewConfUsecase(ldapGroupRepository)

	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	server := grpc.NewServer()
	confgrpc.NewConfServerGrpc(server, usecase)
	fmt.Println("Server Run at :50051")

	err = server.Serve(l)
	if err != nil {
		fmt.Println("Unexpected Error", err)
	}
}
