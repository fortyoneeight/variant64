package errortypes

type Type int64

const (
	None Type = iota
	NotFound
	BadRequest
	InternalError
)

func (t Type) String() string {
	switch t {
	case None:
		return "None"
	case NotFound:
		return "NotFound"
	case BadRequest:
		return "BadRequest"
	case InternalError:
		return "InternalError"
	}
	return ""
}

type TypedError struct {
	errType Type
	errMsg  string
}

func (e TypedError) GetType() Type {
	return e.errType
}

func (e TypedError) Error() string {
	return e.errMsg
}

func (e TypedError) IsError() bool {
	return e.errType == None || len(e.errMsg) > 0
}

func New(t Type, errMsg string) TypedError {
	return TypedError{
		errType: t,
		errMsg:  errMsg,
	}
}
