package model

type NotFoundError struct {
	Err error
}

func (e NotFoundError) Error() string {
	return e.Err.Error()
}

type DuplicateError struct {
	Err error
}

func (e DuplicateError) Error() string {
	return e.Err.Error()
}
