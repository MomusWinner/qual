package models

import (
	"time"
)

type User struct {
	ID        int32
	Name      string
	Email     string
	Password  []byte
	Birthday  *time.Time
	CreatedAt *time.Time
}
