package main

import (
	"fmt"

	"github.com/kw510/go-grpc/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	err := errors.New(codes.DeadlineExceeded, "error message")
	fmt.Println(status.Convert(err))
	// rpc error: code = DeadlineExceeded desc = error message

	err = errors.Errorf(codes.AlreadyExists, "wrap1: %w\n", err)
	fmt.Println(status.Convert(err))
	// rpc error: code = AlreadyExists desc = wrap1: error message

	err = fmt.Errorf("wrap2: %w", err)
	fmt.Println(status.Convert(err))
	// rpc error: code = AlreadyExists desc = wrap2: wrap1: error message

	err = errors.Errorf(codes.FailedPrecondition, "wrap3: %w", err)
	fmt.Println(status.Convert(err))
	// rpc error: code = FailedPrecondition desc = wrap3: wrap2: wrap1: error message

	s := status.New(codes.DeadlineExceeded, "error message")
	fmt.Println(s)
	// rpc error: code = DeadlineExceeded desc = error message

	serr := status.Errorf(codes.AlreadyExists, "wrap1: %v", s.Err())
	fmt.Println(status.Convert(serr))
	// rpc error: code = AlreadyExists desc = wrap1: rpc error: code = DeadlineExceeded desc = error message

	serr = fmt.Errorf("wrap2: %w", serr)
	fmt.Println(status.Convert(serr))
	// rpc error: code = AlreadyExists desc = wrap2: rpc error: code = AlreadyExists desc = wrap1: rpc error: code = DeadlineExceeded desc = error message

	serr = status.Errorf(codes.FailedPrecondition, "wrap3: %v", serr)
	fmt.Println(status.Convert(serr))
	// rpc error: code = FailedPrecondition desc = wrap3: wrap2: rpc error: code = AlreadyExists desc = wrap1: rpc error: code = DeadlineExceeded desc = error message
}
