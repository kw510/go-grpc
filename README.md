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
```go
s := status.New(codes.DeadlineExceeded, "error message")
fmt.Println(s)
// rpc error: code = DeadlineExceeded desc = error message

err := fmt.Errorf("wrap1: %w", s.Err())
fmt.Println(status.Convert(serr))
// rpc error: code = DeadlineExceeded desc = wrap1: rpc error: code = DeadlineExceeded desc = error message

err = fmt.Errorf("wrap2: %w", err)
fmt.Println(status.Convert(serr))
// rpc error: code = DeadlineExceeded desc = wrap2: wrap1: rpc error: code = DeadlineExceeded desc = error message

err = fmt.Errorf("wrap3: %w", err)
fmt.Println(status.Convert(serr))
// rpc error: code = DeadlineExceeded desc = wrap3: wrap2: wrap1: rpc error: code = DeadlineExceeded desc = error message
```