package helper

import (
	"encoding/json"
	"strconv"
)

func StringToPointer(s string) *string {
	return &s
}

func JsonStringify(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func IntToString(i int) string {
	return strconv.Itoa(i)
}
