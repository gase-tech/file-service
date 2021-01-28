package error

// Define your Error struct
type RecordNotFoundError struct {
	msg string
}

// Create a function Error() string and associate it to the struct.
func (error *RecordNotFoundError) Error() string {
	return error.msg
}

// Now you can construct an error object using MyError struct.
func ThisFunctionReturnError(msg string) error {
	return &RecordNotFoundError{msg}
}
