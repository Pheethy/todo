package handler

import (
	"fmt"
	"github/pheethy/todo/helper"
	"github/pheethy/todo/models"
	"github/pheethy/todo/service/todo"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type todoHandler struct {
	todoUs todo.TodoUsecase
}

func NewTodoHandler(todoUs todo.TodoUsecase) todo.TodoHandler {
	return todoHandler{todoUs: todoUs}
}

func (h todoHandler) CreateTask(c *gin.Context) {
	var ctx = c.Request.Context()
	var newTask = new(models.Task)
	var now = helper.NewTimestampFromTime(time.Now())

	if err := c.ShouldBindJSON(newTask); err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("can't binding data: ", err))
		return
	}

	newTask.NewId()
	newTask.SetCreatedAt(now)
	newTask.SetUpatedAt(now)
	newTask.Status = "draft"

	if err := h.todoUs.CreateTask(ctx, newTask); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	resp := map[string]interface{}{
		"message": "Created.",
		"id":      newTask.Id,
	}

	c.JSON(http.StatusOK, resp)
}

func (h todoHandler) FetchListTodo(c *gin.Context) {
	var ctx = c.Request.Context()

	tasks, err := h.todoUs.FetchListTodo(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if len(tasks) < 1 {
		c.JSON(http.StatusNoContent, nil)
		return
	}

	resp := map[string]interface{}{
		"tasks": tasks,
	}

	c.JSON(http.StatusOK, resp)
}
