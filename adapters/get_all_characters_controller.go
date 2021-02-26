package adapters

import (
	"bootcamp/domain/model"
	"bootcamp/usecase"
	"bootcamp/usecase/interfaces"
)

type getAllCharactersController struct {
	UseCase usecase.GetAllCharactersUserCase
}

type GetAllCharactersController interface {
	Execute() (string, error)
}

func NewGetAllCharactersController(repository interfaces.CharacterRepository) GetAllCharactersController {
	controller := &getAllCharactersController{}
	controller.UseCase = usecase.NewGetAllCharactersUseCase(repository)
	return controller
}

func (controller *getAllCharactersController) Execute() (string, error) {
	characters, err := controller.UseCase.Execute()
	if err != nil {
		return "", err
	}
	return controller.Present(characters), nil
}

func (controller *getAllCharactersController) Present(characters []*model.Character) string {
	return "todos"
}
