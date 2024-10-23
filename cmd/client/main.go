package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/notblinkyet/Crud/pkg/config"
	crud "github.com/notblinkyet/crud_gRPC/internal/api/proto"
	"github.com/notblinkyet/crud_gRPC/internal/cli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	config, err := config.ReadConfig("/home/hobonail/go_projects/grpc_project/config.yaml")

	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.NewClient(config.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := crud.NewCrudServiceClient(conn)
	reader := bufio.NewReader(os.Stdin)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		command_list := strings.Split(string(line), " ")
		switch command_list[0] {
		case "exit":
			fmt.Println("Выхожу...")
			return
		default:
			err := cli.RunCLI(command_list, client)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

}
