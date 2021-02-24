package datastore

import (
	"io/ioutil"
)

const characterFile = "characters.csv"

func LoadCharactersFile() ([]byte, error) {
	return ioutil.ReadFile(characterFile)
}
