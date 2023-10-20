package repository

import (
	"context"
	"github/pheethy/todo/helper"
	"github/pheethy/todo/models"
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/BlackMocca/sqlx"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v2"
)

func TestFetchTodo(t *testing.T) {
	now := helper.NewTimestampFromTime(time.Now())
	taskId := uuid.FromStringOrNil("907eefd8-181b-457b-8ca2-692c442b2b0b")
	tasks := []*models.Task{
		&models.Task{
			Id:          &taskId,
			TaskName:    "แก๊งหัวขโมยขนม",
			Status:      "draft",
			CreatorName: "pheethy",
			CreatedAt:   &now,
			UpdatedAt:   &now,
		},
		&models.Task{
			Id:          &taskId,
			TaskName:    "แก๊งหัวขโมยน้ำอัดลม",
			Status:      "draft",
			CreatorName: "pheethy",
			CreatedAt:   &now,
			UpdatedAt:   &now,
		},
	}

	t.Run("success", func(t *testing.T) {
		db, sqlMock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		sqlxDB := sqlx.NewDb(db, "sqlmock")
		defer db.Close()

		rows := sqlmock.NewRows([]string{
			"id", "task_name", "status", "creator_name", "created_at", "updated_at",
		})

		for _, task := range tasks {
			rows.AddRow(
				task.Id.String(), task.TaskName, task.Status, task.CreatorName, task.CreatedAt, task.UpdatedAt,
			)
		}

		sql := `
		SELECT
			(.+)
		FROM
			(.+)
		`

		sqlMock.ExpectQuery(sql).WillReturnRows(rows)

		repo := NewTodoRepository(sqlxDB)
		epTodo, err := repo.FetchListTodo(context.Background())

		assert.NoError(t, err)
		assert.NotEmpty(t, epTodo)
	})
}
