package usecase

import (
	"time"

	"github.com/go-pg/pg/v10/orm"
	"github.com/muhammadikhsaan/simple-api-with-jwt.git/src/helpers/hashing"
	"github.com/muhammadikhsaan/simple-api-with-jwt.git/src/pages/auth/models"
	"github.com/muhammadikhsaan/simple-api-with-jwt.git/src/pages/auth/repository"
)

var (
	registerRepository = repository.NewLoginRepository()
)

func IsRegister(user *models.Users) (orm.Result, error) {
	Password, err := hashing.HashPassword(user.Password)

	if err != nil {
		return nil, err
	}

	user.Password = Password
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	result, err := registerRepository.Post(user)

	if err != nil {
		return nil, err
	}

	return result, nil
}
