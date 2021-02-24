package usecases

import (
	"bootcamp/usecases/interfaces"
)

type getCharacter struct {
	Reposiroty interfaces.CharacterRepository
	Presenter  interfaces.GetCharacterPresenter
}

type GetCharacterUserCase interface {
	Execute(id int) (string, error)
}

func NewGetCharacterUseCase(repository interfaces.CharacterRepository, presenter interfaces.GetCharacterPresenter) GetCharacterUserCase {
	return &getCharacter{repository, presenter}
}

func (usecase *getCharacter) Execute(id int) (string, error) {
	character, err := usecase.Reposiroty.Get(id)
	if err != nil {
		return "", err
	}
	return usecase.Presenter.Present(character), nil
}
