package pkg

// The `import "errors"` statement is importing the `errors` package from the Go standard library. This
// package provides a simple way to create and manipulate errors in Go programs.
import "errors"

// Declaring a variable `ErrUserNotFound` and assigning it a new error instance with the message "user
// not found" using the `errors.New()` function from the `errors` package. This variable can be used to
// represent the specific error of a user not being found in the program.
var (
	ErrUserNotFound = errors.New("user not found")
)
