package errorHandlers

type AlreadyExistsError struct {
	Entity string
}

func (error *AlreadyExistsError) Error() string {
	return error.Entity + " Already Exists."
}
