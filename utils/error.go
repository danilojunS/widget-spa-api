package utils

// CheckError panics if there is an error
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
