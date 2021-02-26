package adapters

import (
	"bootcamp/helper"
	"bootcamp/usecase"
	"bootcamp/usecase/interfaces"
	"fmt"
)

type getCharacterController struct {
	UseCase usecase.GetCharacterUserCase
}

type GetCharacterController interface {
	Execute(id int) (string, error)
}

func NewGetCharacterController(repository interfaces.CharacterRepository) GetCharacterController {
	controller := &getCharacterController{}
	controller.UseCase = usecase.NewGetCharacterUseCase(repository)
	return controller
}

func (controller *getCharacterController) Execute(id int) (string, error) {
	character, err := controller.UseCase.Execute(id)
	if err != nil {
		return "", fmt.Errorf("Error getting the character: %w", err)
	}
	return helper.ToJson(character)
}
