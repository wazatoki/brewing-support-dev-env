package utils

// BoolToInt boolをintに変換する。true -> 1 false -> 0
func BoolToInt(b bool) int {
	if b == true {
		return 1
	}
	return 0
}
