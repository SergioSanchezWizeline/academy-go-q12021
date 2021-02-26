package adapters

import (
	"bootcamp/domain/model"
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
		return "", err
	}
	return controller.Present(character), nil
}

func (controller *getCharacterController) Present(character *model.Character) string {
	return fmt.Sprintf("{ \"id\": %v, \"name\": \"%v\"}", character.Id, character.Name)
}
