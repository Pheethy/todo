package usecase

import (
	"context"
	"github/pheethy/todo/models"
	"github/pheethy/todo/service/todo"
)

type todoUsecase struct {
	todoRepo todo.TodoRepository
}

func NewTodoUsecase(todoRepo todo.TodoRepository) todo.TodoUsecase {
	return todoUsecase{todoRepo: todoRepo}
}

func (u todoUsecase) CreateTask(ctx context.Context, task *models.Task) error {
	return u.todoRepo.CreateTask(ctx, task)
}

func (u todoUsecase) FetchListTodo(ctx context.Context) ([]*models.Task, error) {
	return u.todoRepo.FetchListTodo(ctx)
}
