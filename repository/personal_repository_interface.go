package repository

import "go-api/model"

type PersonalRepositoryInterface interface {
	GetPersonals() ([]model.Personal, error)
	CreatePersonal(personal model.Personal) (int, error)
	UpdatePersonal(personal model.Personal) (model.Personal, error)
	DeletePersonal(id int) error
}
