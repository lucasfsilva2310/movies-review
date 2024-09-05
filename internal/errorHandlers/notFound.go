package errorHandlers

type EntityNotFound struct {
	Entity string
}

func (error *EntityNotFound) Error() string {
	return error.Entity + " Not Found."
}
