package todo

import (
	"context"
	"github/pheethy/todo/models"
)

type TodoRepository interface {
	CreateTask(ctx context.Context, task *models.Task) error
	FetchListTodo(ctx context.Context) ([]*models.Task, error)
}
