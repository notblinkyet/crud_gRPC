package cli

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	crud "github.com/notblinkyet/crud_gRPC/internal/api/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func RunCLI(args []string, client crud.CrudServiceClient) error {
	if len(args) < 2 || args[0] != "my-cli" {
		return errors.New("uknown util")
	}
	switch args[1] {
	case "read":
		if len(args) == 2 {
			tasks, err := client.AllRead(context.Background(), &emptypb.Empty{})
			if err != nil {
				return err
			}
			fmt.Println("Task:")
			for i := range tasks.Tasks {
				fmt.Println(tasks.Tasks[i])
			}
		} else {
			id, err := strconv.Atoi(args[2])
			if err != nil {
				return err
			}

			task, err := client.IdRead(context.Background(), &crud.Id{Id: int64(id)})
			if err != nil {
				return err
			}
			fmt.Printf("%s\n%s\n%s\n\n", task.Title, task.Description, task.Status)

		}

	case "create":
		if len(args) != 5 {
			return errors.New("to create task need title, description and status")
		}
		title, description, status := args[2], args[3], args[4]

		id, err := client.Create(context.Background(), &crud.Task{
			Title:       title,
			Description: description,
			Status:      status,
		})

		if err != nil {
			return err
		}
		fmt.Println(id)

	case "update":
		if len(args) != 6 {
			return errors.New("to create task need title, description and status")
		}
		id, err := strconv.Atoi(args[2])
		if err != nil {
			return err
		}

		title, description, status := args[3], args[4], args[5]

		_, err = client.Update(context.Background(), &crud.UpdateResponse{
			Id: int64(id),
			Task: &crud.Task{
				Title:       title,
				Description: description,
				Status:      status,
			},
		})

		if err != nil {
			return err
		}
		fmt.Printf("Update task with id: %d/n", id)
	case "delete":
		id, err := strconv.Atoi(args[2])
		if err != nil {
			return err
		}

		_, err = client.Delete(context.Background(), &crud.Id{Id: int64(id)})
		if err != nil {
			return err
		}
		fmt.Printf("Delete task with id: %d/n", id)
	default:
		return errors.New("uknown comand;\nneed: create/read/update/delete")
	}
	return nil

}
