package models

import (
	"github/pheethy/todo/helper"

	"github.com/gofrs/uuid"
)

type Task struct {
	Id          *uuid.UUID        `json:"id" db:"id"`
	TaskName    string            `json:"task_name" db:"task_name"`
	Status      string            `json:"status" db:"status"`
	CreatorName string            `json:"creator_name" db:"creator_name"`
	CreatedAt   *helper.Timestamp `json:"created_at" db:"created_at"`
	UpdatedAt   *helper.Timestamp `json:"updated_at" db:"updated_at"`
	DeletedAt   *helper.Timestamp `json:"deleted_at" db:"deleted_at"`
}

func (t *Task) NewId() {
	uid, _ := uuid.NewV4()
	t.Id = &uid
}

func (t *Task) SetCreatedAt(now helper.Timestamp) {
	t.CreatedAt = &now
}

func (t *Task) SetUpatedAt(now helper.Timestamp) {
	t.UpdatedAt = &now
}
