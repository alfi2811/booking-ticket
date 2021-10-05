package business

import "errors"

var (
	ErrInternalServer = errors.New("something gone wrong, contact administrator")

	ErrNotFound = errors.New("data not found")

	ErrDuplicateData = errors.New("duplicate data")

	ErrPasswordWrong = errors.New("password wrong")

	ErrUsernamePasswordNotFound = errors.New("(Username) or (Password) empty")
)
