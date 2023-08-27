package route

import (
	"github/pheethy/todo/service/todo"

	"github.com/gin-gonic/gin"
)

type Route struct {
	e *gin.Engine
}

func NewRoute(e *gin.Engine) *Route {
	return &Route{e: e}
}

func (r Route) RegisterRoute(todoHandle todo.TodoHandler) {
	r.e.POST("/task", todoHandle.CreateTask)
	r.e.GET("/tasks", todoHandle.FetchListTodo)
}
