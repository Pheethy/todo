package models

import (
	"github/pheethy/todo/helper"

	"github.com/gofrs/uuid"
)

type Task struct {
	TableName   struct{}          `json:"-" db:"todo" pk:"ID"`
	Id          *uuid.UUID        `json:"id" db:"id" type:"uuid"`
	TaskName    string            `json:"task_name" db:"task_name" type:"string"`
	Status      string            `json:"status" db:"status" type:"string"`
	CreatorName string            `json:"creator_name" db:"creator_name" type:"string"`
	CreatedAt   *helper.Timestamp `json:"created_at" db:"created_at" type:"timestamp"`
	DeletedAt   *helper.Timestamp `json:"deleted_at" db:"deleted_at" type:"timestamp"`
	UpdatedAt   *helper.Timestamp `json:"updated_at" db:"updated_at" type:"timestamp"`
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
