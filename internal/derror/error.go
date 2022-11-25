package derror

import (
	"fmt"
	"net/http"
)

type ServerError interface {
	error
	Status() int
	Message() string
	Description() string
	TraceMessage() string
	HttpStatusCode() int
	//GrpcStatusCode() codes.Code
}

type (
	errorCode int

	serverError struct {
		code    errorCode
		message string
		desc    string
		trace   string
	}
)

var _ ServerError = (*serverError)(nil)

const (
	invalidArgumentErrorCode errorCode = 1
	unauthorizedErrorCode    errorCode = 2
	notFoundErrorCode        errorCode = 3
	timeoutErrorCode         errorCode = 4
	panicErrorCode           errorCode = 5
	internalErrorCode        errorCode = 6
	unimplementedErrorCode   errorCode = 7
	unknownErrorCode         errorCode = 8
)

var (
	InvalidArgument = &serverError{
		message: "invalid_argument",
		code:    invalidArgumentErrorCode,
	}

	AccessDenied = &serverError{
		message: "access_denied",
		code:    unauthorizedErrorCode,
	}

	NotFound = &serverError{
		message: "not_found",
		code:    notFoundErrorCode,
	}

	Timeout = &serverError{
		message: "timeout",
		code:    timeoutErrorCode,
	}

	Panic = &serverError{
		message: "panic",
		code:    panicErrorCode,
	}

	InternalServer = &serverError{
		message: "internal_server",
		code:    internalErrorCode,
	}

	Unimplemented = &serverError{
		message: "unimplemented",
		code:    unimplementedErrorCode,
	}

	Unknown = &serverError{
		message: "unknown",
		code:    unknownErrorCode,
	}
)

var (
	httpErrors = map[errorCode]int{
		invalidArgumentErrorCode: http.StatusBadRequest,
		unauthorizedErrorCode:    http.StatusUnauthorized,
		notFoundErrorCode:        http.StatusNotFound,
		timeoutErrorCode:         http.StatusRequestTimeout,
		panicErrorCode:           http.StatusInternalServerError,
		internalErrorCode:        http.StatusInternalServerError,
		unimplementedErrorCode:   http.StatusNotImplemented,
		unknownErrorCode:         http.StatusInternalServerError,
	}
)

func (se *serverError) Error() string {
	if len(se.desc) != 0 {
		return fmt.Sprintf("(status_code: %d) %s, desc: %s", se.code, se.message, se.desc)
	}
	return fmt.Sprintf("(status_code: %d) %s", se.code, se.message)
}

func (se *serverError) Status() int {
	return int(se.code)
}

func (se *serverError) Message() string {
	return se.message
}

func (se *serverError) Description() string {
	return se.desc
}

func (se *serverError) TraceMessage() string {
	return se.trace
}

func (se *serverError) HttpStatusCode() int {
	status, ok := httpErrors[se.code]
	if !ok {
		return http.StatusNotFound
	}
	return status
}

