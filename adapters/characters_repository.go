package adapters

import (
	"bootcamp/domain/model"
	"bootcamp/infra/datastore"
	"bootcamp/usecase/interfaces"
	"fmt"
	"strconv"
	"strings"
)

type fileStoreCharacterRepository struct {
	characters map[int]*model.Character
}

func NewCharacterRepository() interfaces.CharacterRepository {
	return &fileStoreCharacterRepository{}
}

func (repository *fileStoreCharacterRepository) load() error {

	if repository.characters != nil {
		return nil
	}

	data, err := datastore.LoadCharactersFile()
	if err != nil {
		return err
	}
	content := string(data)
	lines := strings.Split(content, "\n")

	repository.characters = make(map[int]*model.Character)

	for _, line := range lines {
		parts := strings.Split(line, ",")
		id, err := strconv.Atoi(parts[0])
		if err != nil {
			return err
		}
		character := &model.Character{Id: id, Name: parts[1]}
		repository.characters[character.Id] = character
	}
	return nil
}

func (repository *fileStoreCharacterRepository) All() ([]*model.Character, error) {
	err := repository.load()
	if err != nil {
		return nil, err
	}
	characterList := make([]*model.Character, 0, len(repository.characters))
	for _, character := range repository.characters {
		characterList = append(characterList, character)
	}
	return characterList, nil
}

func (repository *fileStoreCharacterRepository) Get(id int) (*model.Character, error) {
	err := repository.load()
	if err != nil {
		return nil, err
	}
	if character, ok := repository.characters[id]; ok {
		return character, nil
	}
	return nil, fmt.Errorf("Character with id: %v not found", id)
}
