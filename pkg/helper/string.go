package helper

import "encoding/json"

func StringToPointer(s string) *string {
	return &s
}

func JsonStringify(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}
