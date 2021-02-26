package adapters

import (
	"bootcamp/helper"
	"bootcamp/usecase"
	"bootcamp/usecase/interfaces"
	"fmt"
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
		return "", fmt.Errorf("Error getting all characters: %w", err)
	}
	return helper.ToJson(characters)
}
