package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type PersonalUseCase struct {
	repository repository.PersonalRepositoryInterface
}

func NewPersonalUseCase(repo repository.PersonalRepositoryInterface) PersonalUseCase {
	return PersonalUseCase{
		repository: repo,
	}
}

func (uc *PersonalUseCase) GetPersonalByID(id int) (model.Personal, error) {
	return uc.repository.GetPersonalByID(id)
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

func (uc *PersonalUseCase) UpdatePersonal(personal model.Personal) (model.Personal, error) {
	return uc.repository.UpdatePersonal(personal)
}

func (uc *PersonalUseCase) DeletePersonal(id int) error {
	return uc.repository.DeletePersonal(id)
}
