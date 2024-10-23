package main

import (
	"log"
	"net"

	"fmt"

	"github.com/notblinkyet/Crud/pkg/config"
	"github.com/notblinkyet/Crud/pkg/storage/posgresql"
	crud "github.com/notblinkyet/crud_gRPC/internal/api/proto"
	"github.com/notblinkyet/crud_gRPC/internal/server"
	"google.golang.org/grpc"
)

var srv server.GRPCServer

func main() {
	config, err := config.ReadConfig("/home/hobonail/go_projects/grpc_project/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	switch config.DataBase.Type {
	case "postgres":
		connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
			config.DataBase.Username, config.DataBase.Password, config.DataBase.Host, config.DataBase.Port, config.DataBase.Name)

		db, err := posgresql.Open(connString)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			fmt.Println("DataBase connect!")
		}
		srv.Storage = db
	default:
		log.Fatal("unknown BD")
	}

	defer srv.Storage.Close()

	s := grpc.NewServer()
	crud.RegisterCrudServiceServer(s, srv)

	lis, err := net.Listen("tcp", config.Addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(s.Serve(lis))
}
