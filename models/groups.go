package models

import "github.com/google/uuid"

type Group struct {
	ID        uuid.UUID `json:"id" sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name      string    `json:"name"`
	CreatedAt int64     `json:"created_at" gorm:"autoCreateTime" gorm:"<-:create"`
	UpdatedAt int64     `json:"updated_at" gorm:"autoUpdateTime"`
}
