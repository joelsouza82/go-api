package usecase

import "go-api/model"

type LoginUseCaseInterface interface {
	GetLogins() ([]model.Login, error)
	GetLoginByID(id int) (model.Login, error)
	CreateLogin(login model.Login) (model.Login, error)
	UpdateLogin(login model.Login) (model.Login, error)
	DeleteLogin(id int) error
}
