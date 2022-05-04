package main

import (
	"fmt"
	"github.com/samithiwat/samithiwat-backend-user/src/config"
	"github.com/samithiwat/samithiwat-backend-user/src/database"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Cannot load config", err.Error())
	}

	_, err = database.InitDatabase(&conf.Database)
	if err != nil {
		log.Fatal("Cannot connect to database: ", err.Error())
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", conf.App.Port))

	grpcServer := grpc.NewServer()

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}

	fmt.Println(fmt.Sprintf("fe camp user service listening at port %v", conf.App.Port))
}
