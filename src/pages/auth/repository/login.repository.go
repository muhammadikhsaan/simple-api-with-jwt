package repository

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/muhammadikhsaan/simple-api-with-jwt.git/src/config"
	"github.com/muhammadikhsaan/simple-api-with-jwt.git/src/pages/auth/models"
)

type LoginRepository struct {
	db *pg.DB
}

func NewLoginRepository() AuthRepository {
	return &LoginRepository{
		db: config.DB,
	}
}

func (lg *LoginRepository) Get(user *models.Users) error {
	err := lg.db.Model(user).Where("email ilike ?", user.Email).Select()

	return err
}

func (lg *LoginRepository) Post(user *models.Users) (orm.Result, error) {
	result, err := lg.db.Model(user).Insert()

	return result, err
}
