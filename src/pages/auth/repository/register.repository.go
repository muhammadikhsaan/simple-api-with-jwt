package repository

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/muhammadikhsaan/simple-api-with-jwt.git/src/config"
	"github.com/muhammadikhsaan/simple-api-with-jwt.git/src/pages/auth/models"
)

type RegisterRepository struct {
	db *pg.DB
}

func NewRegisterRepository() AuthRepository {
	return &LoginRepository{
		db: config.DB,
	}
}

func (lg *RegisterRepository) Post(user *models.Users) (orm.Result, error) {
	result, err := lg.db.Model(user).Insert()

	return result, err
}
