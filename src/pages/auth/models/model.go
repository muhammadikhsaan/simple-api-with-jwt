package models

import "time"

type Users struct {
	Email     string    `db:"email" json:"email" validate:"required,email"`
	Password  string    `db:"password" json:"password" validate:"required"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
}
