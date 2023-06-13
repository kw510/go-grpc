# Go gRPC errors

A wrapper libary. Wraps errors with gRPC statuses. This allows for this sort of use:

```go
err := errors.New(codes.DeadlineExceeded, "error message")
fmt.Println(status.Convert(err))
// rpc error: code = DeadlineExceeded desc = error message

err = errors.Errorf(codes.AlreadyExists, "wrap1: %w", err)
fmt.Println(status.Convert(err))
// rpc error: code = AlreadyExists desc = wrap1: error message

err = fmt.Errorf("wrap2: %w", err)
fmt.Println(status.Convert(err))
// rpc error: code = AlreadyExists desc = wrap2: wrap1: error message

err = errors.Errorf(codes.FailedPrecondition, "wrap3: %w", err)
fmt.Println(status.Convert(err))
// rpc error: code = FailedPrecondition desc = wrap3: wrap2: wrap1: error message
```

# Why
As of current, gRPC status codes are not carried via the `fmt.Errorf` function, nor can you override the code later. If you try to do it as of now, using the sample above, this is the best you can achieve:

Update: I found out there are carried via `fmt.Errorf` function :tada: This wrapper lib will just remove the code from the error message to make it look pretty. It loses the history of the code, but at least its nice to look at :wink:
```go
s := status.New(codes.DeadlineExceeded, "error message")
fmt.Println(s)
// rpc error: code = DeadlineExceeded desc = error message

err := status.Errorf(codes.AlreadyExists, "wrap1: %v", s.Err())
fmt.Println(status.Convert(err))
// rpc error: code = AlreadyExists desc = wrap1: rpc error: code = DeadlineExceeded desc = error message

err = fmt.Errorf("wrap2: %w", err)
fmt.Println(status.Convert(err))
// rpc error: code = AlreadyExists desc = wrap2: rpc error: code = AlreadyExists desc = wrap1: rpc error: code = DeadlineExceeded desc = error message

err = status.Errorf(codes.FailedPrecondition, "wrap3: %v", serr)
fmt.Println(status.Convert(err))
// rpc error: code = FailedPrecondition desc = wrap3: wrap2: rpc error: code = AlreadyExists desc = wrap1: rpc error: code = DeadlineExceeded desc = error message
```