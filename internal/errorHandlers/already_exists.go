package errorHandlers

type AlreadyExistsError struct {
	Entity string
}

func (error *AlreadyExistsError) Error() string {
	return error.Entity + " Already Exists."
}

type EntityNotFound struct {
	Entity string
}

func (error *EntityNotFound) Error() string {
	return error.Entity + " Not Found."
}
