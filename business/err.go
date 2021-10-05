package business

import "errors"

var (
	ErrEmailorPass    = errors.New("email or password is incorrect")
	ErrDuplicateData  = errors.New("account already exist")
	ErrInternalServer = errors.New("internal server error")
)
