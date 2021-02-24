package interfaces

import (
	"bootcamp/domain/model"
)

type GetCharacterPresenter interface {
	Present(character *model.Character) string
}
