package error

const (
	TypeValidation      ErrorType = "VALIDATION"
	TypeNotFound                  = "NOT_FOUND"
	TypeUnAuthorized              = "UNAUTHORIZED"
	TypeForbidden                 = "FORBIDDEN"
	TypeInternal                  = "INTERNAL"
	TypeDuplicate                 = "DUPLICATE"
	TypeTooManyRequests           = "TOO_MANY_REQUESTS"
)

type ErrorType string

func (error *err) Is(errType ErrorType) bool {
	return error.Type() == errType
}

func Is(err error, errType ErrorType) bool {
	if e, ok := err.(ErrorModel); ok {
		return e.Is(errType)
	}
	return false
}
