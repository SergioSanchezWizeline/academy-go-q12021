package interfaces

import (
	"bootcamp/domain/model"
)

type CharacterRepository interface {
	All() ([]*model.Character, error)
	Get(id int) (*model.Character, error)
}
