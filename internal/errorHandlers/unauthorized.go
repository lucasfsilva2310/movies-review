package errorHandlers

type Unauthorized struct {
}

func (error *Unauthorized) Error() string {
	return "Not authorized."
}
