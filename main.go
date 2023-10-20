package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github/pheethy/todo/config"
	"github/pheethy/todo/migration/database"
	"github/pheethy/todo/route"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github/pheethy/todo/service/todo/handler"
	"github/pheethy/todo/service/todo/repository"
	"github/pheethy/todo/service/todo/usecase"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}

func main() {
	var ctx = context.Background()
	var cfg = config.LoadConfig(envPath())
	var psqlDB = database.DBConnect(ctx, cfg.Db())
	defer psqlDB.Close()

	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	r := gin.Default()

	todoRepo := repository.NewTodoRepository(psqlDB)
	todoUs := usecase.NewTodoUsecase(todoRepo)
	todoHand := handler.NewTodoHandler(todoUs)
	route := route.NewRoute(r)
	route.RegisterRoute(todoHand)

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Hello web")
	})

	s := &http.Server{
		Addr:           cfg.App().Url(),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err)
		}
	}()
	<- ctx.Done()
	stop()
	fmt.Println("shutting down gracefully.")

	timeOutCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.Shutdown(timeOutCtx); err != nil {
		log.Fatal(err)
	}

}
