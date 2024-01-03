package helper

import "strconv"

func ToDbId(idStr string) int {
	id, _ := strconv.Atoi(idStr)
	return id
}
