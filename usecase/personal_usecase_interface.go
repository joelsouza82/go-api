package usecase

import "go-api/model"

type PersonalUseCaseInterface interface {
	GetPersonalByID(id int) (model.Personal, error)
	GetPersonals() ([]model.Personal, error)
	CreatePersonal(personal model.Personal) (model.Personal, error)
	UpdatePersonal(personal model.Personal) (model.Personal, error)
	DeletePersonal(id int) error
}
