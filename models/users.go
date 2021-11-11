package models

import "github.com/google/uuid"

type Users struct {
	ID        uuid.UUID `json:"id" sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
	CreatedAt int64     `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int64     `json:"updated_at" gorm:"autoUpdateTime"`
}
