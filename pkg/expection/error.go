package expection

type Exception struct {
	Code int
	Message string
}

func (e *Exception) Error() string {
	return e.Message
}

