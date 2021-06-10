package repository

import (
	"github.com/go-pg/pg/v10/orm"
	"github.com/muhammadikhsaan/simple-api-with-jwt.git/src/pages/auth/models"
)

type AuthRepository interface {
	Get(*models.Users) error
	Post(*models.Users) (orm.Result, error)
}
