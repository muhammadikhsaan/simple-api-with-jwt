package usecase

import (
	"errors"

	"github.com/muhammadikhsaan/simple-api-with-jwt.git/src/helpers/hashing"
	"github.com/muhammadikhsaan/simple-api-with-jwt.git/src/pages/auth/models"
	"github.com/muhammadikhsaan/simple-api-with-jwt.git/src/pages/auth/repository"
)

var (
	loginRepository = repository.NewLoginRepository()
)

func IsLogin(user *models.Users) (bool, error) {
	req := *user

	err := loginRepository.Get(user)

	if err != nil || !hashing.CheckPasswordHash(req.Password, user.Password) {
		return false, errors.New("credential not valid")
	}

	return true, nil
}
