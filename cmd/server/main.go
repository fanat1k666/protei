package main

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net"
	"net/http"
	"protei/internal/ems"
	"protei/internal/server"
	proteigrpc "protei/internal/user"

	"google.golang.org/grpc"
)

func main() {
	if err := initCfg(); err != nil {
		log.Fatalf("error initializing congigs: %s", err.Error())
	}
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", viper.GetString("server.port")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	e := ems.NewClient(&http.Client{},
		"https://"+viper.GetString("crdntls.address")+":"+viper.GetString("crdntls.port"),
		viper.GetString("crdntls.login"),
		viper.GetString("crdntls.password"))
	s := grpc.NewServer()
	proteigrpc.RegisterUserSearchServer(s, server.NewServer(e))
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func initCfg() error {
	viper.AddConfigPath("cfg")
	viper.SetConfigName("cfg")
	viper.SetConfigType("json")
	return viper.ReadInConfig()
}
