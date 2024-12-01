package helper

import "reflect"

func Contains(slice interface{}, value interface{}) bool {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return false
	}

	for i := 0; i < v.Len(); i++ {
		if v.Index(i).Interface() == value {
			return true
		}
	}

	return false
}
