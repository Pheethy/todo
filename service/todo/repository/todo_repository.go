package repository

import (
	"context"
	"errors"
	"github/pheethy/todo/constants"
	"github/pheethy/todo/models"
	"github/pheethy/todo/service/todo"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
)

type todoRepository struct {
	db *sqlx.DB
}

func NewTodoRepository(db *sqlx.DB) todo.TodoRepository {
	return todoRepository{db: db}
}

func (t todoRepository) CreateTask(ctx context.Context, task *models.Task) error {
	tx, err := t.db.Beginx()
	if err != nil {
		panic(err)
	}

	sql := `
		INSERT INTO todo (
			id,
			task_name,
			status,
			creator_name,
			created_at,
			updated_at,
			deleted_at
		)
		VALUES(
			$1::uuid,
			$2::text,
			$3::todo_status,
			$4::text,
			$5::timestamp,
			$6::timestamp,
			$7::timestamp
		)
	`
	stmt, err := tx.PreparexContext(ctx, sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.ExecContext(ctx,
		task.Id,
		task.TaskName,
		task.Status,
		task.CreatorName,
		task.CreatedAt,
		task.UpdatedAt,
		task.DeletedAt,
	); err != nil {
		if strings.Contains(err.Error(), constants.ERROR_TASKNAME_WAS_DUPLICATE) {
			return errors.New(constants.ERROR_TASKNAME_WAS_DUPLICATE_SERVICE)
		}
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (t todoRepository) FetchListTodo(ctx context.Context) ([]*models.Task, error) {
	sql := `
	SELECT
		id,
		task_name,
		status,
		creator_name,
		created_at,
		updated_at
	FROM
		todo
	`
	rows, err := t.db.QueryxContext(ctx, sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	return t.orm(rows)
}

func (t todoRepository) orm(rows *sqlx.Rows) ([]*models.Task, error) {
	var tasks = make([]*models.Task, 0)

	for rows.Next() {
		task := &models.Task{}
		err := rows.Scan(
			&task.Id,
			&task.TaskName,
			&task.Status,
			&task.CreatorName,
			&task.CreatedAt, // Scan as string
			&task.UpdatedAt, // Scan as string
		)
		if err != nil {
			log.Println("err here")
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
