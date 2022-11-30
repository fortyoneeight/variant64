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

type TypedError interface {
	GetType() Type
	Error() string
}
