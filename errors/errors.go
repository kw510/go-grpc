package errors

import (
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Error struct {
	codes.Code
	err wrapError
}

func New(code codes.Code, text string) error {
	err := errors.New(text)
	return &Error{
		Code: code,
		err:  wrapError{err},
	}
}

func Errorf(code codes.Code, format string, a ...any) error {
	err := fmt.Errorf(format, a...)
	return Error{
		Code: code,
		err:  wrapError{err},
	}
}

func (e Error) Error() string {
	return e.err.Error()
}

func (e Error) GRPCStatus() *status.Status {
	return status.New(e.Code, e.Error())
}

func (e Error) Unwrap() error {
	return errors.Unwrap(e.err)
}

type wrapError struct {
	err error
}

func (e wrapError) Error() string {
	return e.err.Error()
}

func (e wrapError) GRPCStatus() *status.Status {
	type grpcstatus interface{ GRPCStatus() *status.Status }
	if gs, ok := e.err.(grpcstatus); ok {
		return status.New(gs.GRPCStatus().Code(), e.Error())
	}

	if gs, ok := e.Unwrap().(grpcstatus); ok {
		return status.New(gs.GRPCStatus().Code(), e.Error())
	}

	return status.New(codes.Unknown, e.Error())
}

func (e wrapError) Unwrap() error {
	return e.err
}
