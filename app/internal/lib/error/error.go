package error

func (e ErrInvalidField) Error() string {
	return "Field name is empty. Name='" + string(e) + "'"
}
