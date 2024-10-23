package server

import (
	"context"

	"github.com/notblinkyet/Crud/pkg/storage"
	crud "github.com/notblinkyet/crud_gRPC/internal/api/proto"
	"github.com/notblinkyet/crud_gRPC/internal/convert"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCServer struct {
	Storage storage.Storage
	crud.UnimplementedCrudServiceServer
}

// AllRead implements crud.CrudServiceServer.
func (server GRPCServer) AllRead(ctx context.Context, in *emptypb.Empty) (*crud.Tasks, error) {
	unparce_tasks, err := server.Storage.ReadAll()
	tasks := make([]*crud.Task, len(unparce_tasks))

	if err != nil {
		return nil, err
	}
	for i, unparce_task := range unparce_tasks {
		tasks[i] = convert.ModelToPro(&unparce_task)
	}
	return &crud.Tasks{Tasks: tasks}, err
}

// Create implements crud.CrudServiceServer.
func (server GRPCServer) Create(ctx context.Context, in *crud.Task) (*crud.Id, error) {
	id, err := server.Storage.Create(
		convert.ProToModel(in),
	)
	return &crud.Id{Id: int64(id)}, err
}

// Delete implements crud.CrudServiceServer.
func (server GRPCServer) Delete(ctx context.Context, in *crud.Id) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, server.Storage.Delete(int(in.Id))
}

// IdRead implements crud.CrudServiceServer.
func (server GRPCServer) IdRead(ctx context.Context, in *crud.Id) (*crud.Task, error) {
	unparce_task, err := server.Storage.ReadId(int(in.Id))

	return convert.ModelToPro(unparce_task), err
}

// Update implements crud.CrudServiceServer.
func (server GRPCServer) Update(ctx context.Context, in *crud.UpdateResponse) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, server.Storage.Update(int(in.Id), in.Task.Title, in.Task.Description, in.Task.Status)
}
