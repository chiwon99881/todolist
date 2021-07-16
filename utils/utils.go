package utils

// HandleError is handling error function
func HandleError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
