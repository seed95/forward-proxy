package derror

type (
	errorBuilder struct {
		code    errorCode
		desc    string
		message string
		trace   string
	}

	ErrorBuilder interface {
		Desc(desc string) ErrorBuilder
		Trace(trace string) ErrorBuilder
		Build() ServerError
	}
)

var _ ErrorBuilder = (*errorBuilder)(nil)

var (
	InvalidArgumentBuilder = &errorBuilder{
		message: InvalidArgument.message,
		code:    InvalidArgument.code,
	}

	AccessDeniedBuilder = errorBuilder{
		message: AccessDenied.message,
		code:    AccessDenied.code,
	}

	NotFoundBuilder = &errorBuilder{
		message: NotFound.message,
		code:    NotFound.code,
	}

	TimeoutBuilder = &errorBuilder{
		message: Timeout.message,
		code:    Timeout.code,
	}

	PanicBuilder = &errorBuilder{
		message: Panic.message,
		code:    Panic.code,
	}

	InternalServerBuilder = &errorBuilder{
		message: InternalServer.message,
		code:    InternalServer.code,
	}

	UnimplementedBuilder = &errorBuilder{
		message: Unimplemented.message,
		code:    Unimplemented.code,
	}

	UnknownBuilder = &errorBuilder{
		code:    Unknown.code,
		message: Unknown.message,
	}
)

func (b *errorBuilder) Desc(desc string) ErrorBuilder {
	b.desc = desc
	return b
}

func (b *errorBuilder) Trace(trace string) ErrorBuilder {
	b.trace = trace
	return b
}

func (b *errorBuilder) Build() ServerError {
	return &serverError{
		message: b.message,
		code:    b.code,
		desc:    b.desc,
		trace:   b.trace,
	}
}
