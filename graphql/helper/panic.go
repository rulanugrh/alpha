package helper

func PanicIfError(err error) error {
	if err != nil {
		panic(err)
	}
	return nil
}
