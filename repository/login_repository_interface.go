package repository

import "go-api/model"

type LoginRepositoryInterface interface {
	GetLogins() ([]model.Login, error)
	GetLoginByID(id int) (model.Login, error)
	CreateLogin(login model.Login) (int, error)
	UpdateLogin(login model.Login) (model.Login, error)
	DeleteLogin(id int) error
}
