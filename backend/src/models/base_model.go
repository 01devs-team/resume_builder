package models

import "github.com/google/uuid"

type Base struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
}
