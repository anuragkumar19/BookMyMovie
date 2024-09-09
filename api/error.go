package api

import (
	"errors"

	"bookmymovie.app/bookmymovie/services/serviceserrors"
	"connectrpc.com/connect"
)

var errInternal = errors.New("internal server error")

func serviceErrorHandler(err error) *connect.Error {
	sErr := new(serviceserrors.Error)

	if ok := errors.As(err, &sErr); !ok {
		// TODO: log error
		return connect.NewError(connect.CodeInternal, errInternal)
	}

	switch sErr.Type() {
	case serviceserrors.ErrorTypeInvalidArgument:
		return connect.NewError(connect.CodeInvalidArgument, err)
	case serviceserrors.ErrorTypeNotFound:
		return connect.NewError(connect.CodeNotFound, err)
	case serviceserrors.ErrorTypeAlreadyExist:
		return connect.NewError(connect.CodeAlreadyExists, err)
	case serviceserrors.ErrorTypePermissionDenied:
		return connect.NewError(connect.CodePermissionDenied, err)
	case serviceserrors.ErrorResourceExhausted:
		return connect.NewError(connect.CodeResourceExhausted, err)
	case serviceserrors.ErrorConflict:
		return connect.NewError(connect.CodeAborted, err)
	case serviceserrors.ErrorUnauthenticated:
		return connect.NewError(connect.CodeUnauthenticated, err)
	default:
		return connect.NewError(connect.CodeUnknown, err)
	}
}
