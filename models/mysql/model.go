package mysql

import (
	"time"
)

type CommonColumn struct {
	Id uint64 `json:"id" gorm:"primaryKey"`
	// DataJson  string    `json:"data_json"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
