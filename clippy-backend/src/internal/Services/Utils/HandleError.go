package Utils

func HandleError(err error) {
	if err != nil {
			panic(err)
	}
}