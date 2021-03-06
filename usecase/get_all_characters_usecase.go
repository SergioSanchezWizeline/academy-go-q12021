package usecase

import (
	"bootcamp/domain/model"
	"bootcamp/usecase/interfaces"
)

type getAllCharacters struct {
	Repository interfaces.CharacterRepository
}

type GetAllCharactersUserCase interface {
	Execute() ([]*model.Character, error)
}

func NewGetAllCharactersUseCase(repository interfaces.CharacterRepository) GetAllCharactersUserCase {
	return &getAllCharacters{repository}
}

func (usecase *getAllCharacters) Execute() ([]*model.Character, error) {
	characters, err := usecase.Repository.All()
	if err != nil {
		return nil, err
	}
	return characters, nil
}
