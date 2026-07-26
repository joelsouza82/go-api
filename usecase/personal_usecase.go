package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type PersonalUseCase struct {
	repository repository.PersonalRepository
}

func NewPersonalUseCase(repo repository.PersonalRepository) PersonalUseCase {
	return PersonalUseCase{
		repository: repo,
	}
}

func (uc *PersonalUseCase) GetPersonals() ([]model.Personal, error) {
	return uc.repository.GetPersonals()
}

func (uc *PersonalUseCase) CreatePersonal(personal model.Personal) (model.Personal, error) {
	personalId, err := uc.repository.CreatePersonal(personal)
	if err != nil {
		return model.Personal{}, err
	}

	personal.ID = personalId
	return personal, nil
}
