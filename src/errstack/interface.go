package errstack

// ErrStack represents a 'stack' of errors
type ErrStack interface {
	// implement the error interface
	error

	// Root returns the original error
	Root() error

	// Stack returns a slice of strings that are in the order of the call stack.
	// The final string is the output of Error() from the root error
	Stack() []string

	// Join returns a string made by joining the 
	Join(sep string) string
}