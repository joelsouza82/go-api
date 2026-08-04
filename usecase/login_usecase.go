package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type LoginUseCase struct {
	repository repository.LoginRepositoryInterface
}

func NewLoginUseCase(repo repository.LoginRepositoryInterface) *LoginUseCase {
	return &LoginUseCase{
		repository: repo,
	}
}

func (uc *LoginUseCase) GetLogins() ([]model.Login, error) {
	return uc.repository.GetLogins()
}

func (uc *LoginUseCase) GetLoginByID(id int) (model.Login, error) {
	return uc.repository.GetLoginByID(id)
}

func (uc *LoginUseCase) CreateLogin(login model.Login) (model.Login, error) {
	loginId, err := uc.repository.CreateLogin(login)
	if err != nil {
		return model.Login{}, err
	}

	login.ID = loginId
	return login, nil
}

func (uc *LoginUseCase) UpdateLogin(login model.Login) (model.Login, error) {
	return uc.repository.UpdateLogin(login)
}

func (uc *LoginUseCase) DeleteLogin(id int) error {
	return uc.repository.DeleteLogin(id)
}
