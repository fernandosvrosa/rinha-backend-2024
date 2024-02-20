package error

type InsufficientFund struct {
	Message string
}

func (e InsufficientFund) Error() string {
	return e.Message
}
