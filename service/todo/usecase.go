package todo

import (
	"context"
	"github/pheethy/todo/models"
)

type TodoUsecase interface {
	CreateTask(ctx context.Context, task *models.Task) error
	FetchListTodo(ctx context.Context) ([]*models.Task, error)
}
