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
