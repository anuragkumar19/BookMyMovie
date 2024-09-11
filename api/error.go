package api

import (
	"errors"

	"bookmymovie.app/bookmymovie/services"
	"connectrpc.com/connect"
)

var errInternal = errors.New("internal server error")

func serviceErrorHandler(err error) *connect.Error {
	sErr := new(services.Error)

	if ok := errors.As(err, &sErr); !ok {
		// TODO: log error
		return connect.NewError(connect.CodeInternal, errInternal)
	}

	switch sErr.Type() {
	case services.ErrorTypeInvalidArgument:
		return connect.NewError(connect.CodeInvalidArgument, err)
	case services.ErrorTypeNotFound:
		return connect.NewError(connect.CodeNotFound, err)
	case services.ErrorTypeAlreadyExist:
		return connect.NewError(connect.CodeAlreadyExists, err)
	case services.ErrorTypePermissionDenied:
		return connect.NewError(connect.CodePermissionDenied, err)
	case services.ErrorTypeResourceExhausted:
		return connect.NewError(connect.CodeResourceExhausted, err)
	case services.ErrorTypeConflict:
		return connect.NewError(connect.CodeAborted, err)
	case services.ErrorTypeUnauthenticated:
		return connect.NewError(connect.CodeUnauthenticated, err)
	default:
		return connect.NewError(connect.CodeUnknown, err)
	}
}
