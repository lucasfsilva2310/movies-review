package errorHandlers

type NotSameUser struct {
	Entity string
}

func (error *NotSameUser) Error() string {
	return "Can not look for data from differente user."
}
