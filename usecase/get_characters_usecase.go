package usecase

import (
	"bootcamp/domain/model"
	"bootcamp/usecase/interfaces"
)

type getCharacter struct {
	Repository interfaces.CharacterRepository
}

type GetCharacterUserCase interface {
	Execute(id int) (*model.Character, error)
}

func NewGetCharacterUseCase(repository interfaces.CharacterRepository) GetCharacterUserCase {
	return &getCharacter{repository}
}

func (usecase *getCharacter) Execute(id int) (*model.Character, error) {
	character, err := usecase.Repository.Get(id)
	if err != nil {
		return nil, err
	}
	return character, nil
}
