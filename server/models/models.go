package models

import (
	"time"
)

type KeyValue struct {
	UserID    int       `json:"user_id" gorm:"primaryKey"`
	AppName   string    `json:"app_name" gorm:"primaryKey"`
	Key       string    `json:"key" gorm:"primaryKey"`
	Value     string    `json:"value"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type APIKey struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	UserID    int        `json:"user_id"`
	AppName   string     `json:"app_name"`
	Key       string     `json:"key" gorm:"uniqueIndex"`
	Name      string     `json:"name"`
	IsActive  bool       `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time  `json:"created_at"`
	LastUsed  *time.Time `json:"last_used,omitempty"`
}
