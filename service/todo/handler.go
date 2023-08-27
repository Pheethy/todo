package todo

import (
	"github.com/gin-gonic/gin"
)

type TodoHandler interface {
	CreateTask(c *gin.Context)
	FetchListTodo(c *gin.Context)
}