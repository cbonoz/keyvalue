package models

import (
	"time"

	"github.com/google/uuid"
)

type App struct {
	ID              uint       `json:"id" gorm:"primaryKey"`
	CreatedByUserID uuid.UUID  `json:"created_by_user_id"`
	Name            string     `json:"name"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty"`
}

type KeyValue struct {
	AppID     int32 `json:"app_id"`
	Key       string `json:"key"`
	Value     string `json:"value"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type APIKey struct {
	ID        int32       `json:"id"`
	AppID     int32     `json:"app_id"`
	Key       string     `json:"key"`
	CreatedAt time.Time  `json:"created_at"`
	LastUsed  *time.Time `json:"last_used,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
