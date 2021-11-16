package Utils

import (
	"strconv"
)

func IntToString(any int) string {
	return strconv.Itoa(any)
}

func StringToInt(any string) (int, error) {
	return strconv.Atoi(any)
}
