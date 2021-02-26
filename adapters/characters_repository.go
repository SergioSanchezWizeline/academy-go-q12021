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
	characters    []*model.Character
	charactersMap map[int]*model.Character
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

	repository.characters = make([]*model.Character, len(lines))
	repository.charactersMap = make(map[int]*model.Character)

	for index, line := range lines {
		parts := strings.Split(line, ",")
		id, err := strconv.Atoi(parts[0])
		if err != nil {
			return err
		}
		character := &model.Character{Id: id, Name: parts[1]}
		repository.characters[index] = character
		repository.charactersMap[character.Id] = character
	}
	return nil
}

func (repository *fileStoreCharacterRepository) All() ([]*model.Character, error) {
	err := repository.load()
	if err != nil {
		return nil, err
	}
	return repository.characters, nil
}

func (repository *fileStoreCharacterRepository) Get(id int) (*model.Character, error) {
	err := repository.load()
	if err != nil {
		return nil, err
	}
	if character, ok := repository.charactersMap[id]; ok {
		return character, nil
	}
	return nil, fmt.Errorf("Character with id: %v not found", id)
}
