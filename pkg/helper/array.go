package helper

func Contains(slice []interface{}, value interface{}) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
