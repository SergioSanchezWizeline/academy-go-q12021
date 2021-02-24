package adapters

import (
	"bootcamp/domain/model"
	"bootcamp/usecases"
	"bootcamp/usecases/interfaces"
	"fmt"
)

type getCharacterController struct {
	UseCase usecases.GetCharacterUserCase
}

type GetCharacterController interface {
	Execute(id int) (string, error)
	Present(character *model.Character) string
}

func NewGetCharacterController(repository interfaces.CharacterRepository) GetCharacterController {
	controller := &getCharacterController{}
	controller.UseCase = usecases.NewGetCharacterUseCase(repository, controller)
	return controller
}

func (controller *getCharacterController) Execute(id int) (string, error) {
	return controller.UseCase.Execute(id)
}

func (controller *getCharacterController) Present(character *model.Character) string {
	return fmt.Sprintf("{ \"id\": %v, \"name\": \"%v\"}", character.Id, character.Name)
}
