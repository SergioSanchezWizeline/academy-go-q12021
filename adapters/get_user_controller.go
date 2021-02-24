package adapters

import "bootcamp/usecases"

type getCharacterController struct {
	UseCase usecases.GetCharacterUserCase
}

type GetCharacterController interface {
	Execute(id int) (string, error)
}

func NewGetCharacterController(usecase usecases.GetCharacterUserCase) GetCharacterController {
	return &getCharacterController{usecase}
}

func (controller *getCharacterController) Execute(id int) (string, error) {
	return controller.UseCase.Execute(id)
}
