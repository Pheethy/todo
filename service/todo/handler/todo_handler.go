package handler

import (
	"github/pheethy/todo/helper"
	"github/pheethy/todo/models"
	"github/pheethy/todo/service/todo"
	"log"
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

	err := c.ShouldBindJSON(newTask)
	if err != nil {
		log.Println(err)
	}

	newTask.NewId()
	newTask.SetCreatedAt(now)
	newTask.SetUpatedAt(now)
	newTask.Status = "draft"

	err = h.todoUs.CreateTask(ctx, newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
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

	resp := map[string]interface{}{
		"tasks": tasks,
	}

	c.JSON(http.StatusOK, resp)
}
