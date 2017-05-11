package exoerror

import "fmt"

// Error : Error struct.
type Error struct {
	Code      int
	Message   string
	Retriable bool
}

// Error constants
var (
	ClientError   = &Error{Code: 0, Message: "Failed while initializing client"}
	RequestError  = &Error{Code: 1, Message: "Failed while initializing client"}
	ResponseError = &Error{Code: 2, Message: "Failed generating Exotel response"}
	BadRequest    = &Error{Code: 3, Message: "Bad Request"}
	MissingParams = &Error{Code: 4, Message: "Missing mandatory parameters"}
	CallFailed    = &Error{Code: 4, Message: "Call Failed"}
	DecodeError   = &Error{Code: 5, Message: "Error while decoding response"}
	AuthError     = &Error{Code: 6, Message: "Error setting auth"}
	RequestFailed = &Error{Code: 7, Message: "Failed while making request"}
)

func (e *Error) Error() string {
	if e == nil {
		return fmt.Sprint("[EXOERROR] Code:0;\tMessage:Invalid error")
	}
	return fmt.Sprintf("[EXOERROR] Code:%d;\tMessage:%s", e.Code, e.Message)
}

//String implements fmt.Stringer
func (e *Error) String() string {
	return e.Error()
}

//New returns a new instance of Error with the
func New(s interface{}) (err *Error) {
	switch s := s.(type) {
	default:
		err = new(Error)
	case Error:
		err = new(Error)
		*err = s
	case *Error:
		err = s
	case error:
		err = new(Error)
		err.Message = s.Error()
		err.Code = 100
	case string:
		err = new(Error)
		err.Message = s
		err.Code = 100
	}
	return
}
