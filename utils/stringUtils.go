package utils

//dosya icindeki tum stringi data olarak ele aldık.
func IsEmpty(data string) bool {
	if len(data) == 0 {
		return true
	}
	return false
}
