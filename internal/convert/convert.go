package convert

import (
	"github.com/notblinkyet/Crud/pkg/models"
	crud "github.com/notblinkyet/crud_gRPC/internal/api/proto"
)

func ProToModel(task *crud.Task) *models.Task {
	return &models.Task{
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
	}
}

func ModelToPro(task *models.Task) *crud.Task {
	return &crud.Task{
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
	}
}
