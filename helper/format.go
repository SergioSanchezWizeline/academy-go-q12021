package helper

import (
	"encoding/json"
	"fmt"
)

func ToJson(entity interface{}) (result string, err error) {
	data, err := json.Marshal(entity)
	if err != nil {
		err = fmt.Errorf("Json serialization error: %w", err)
		return
	}
	result = string(data)
	return
}
