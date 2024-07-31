package state

type StateError struct {
	message string
}

func (e StateError) Error() string {
	return e.message
}

func newStateError(message string) StateError {
	return StateError{message: message}
}
