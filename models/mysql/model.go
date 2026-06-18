package mysql

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

var (
	ErrNoFieldsToUpdate = errors.New("no fields to update")
)

type IDColumn struct {
	ID uint64 `json:"id" gorm:"primaryKey"`
}

type CommonColumn struct {
	ID        uint64         `json:"id" gorm:"primaryKey"`
	DataJson  *string        `json:"data_json" gorm:"column:data_json"`
	CreatedAt *time.Time     `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt *time.Time     `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at;index"`
}
