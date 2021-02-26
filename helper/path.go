package helper

import (
	"strconv"
)

func ExtractId(prefix, path string) (id int, err error) {
	idStr := path[len(prefix):]
	id, err = strconv.Atoi(idStr)
	return
}
