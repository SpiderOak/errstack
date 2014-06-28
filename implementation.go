package errstack

import (
	"fmt"
	"strings"
)

const (
	defaultSeparator = "|"
)

type errorStack struct {

	// root is the original error
	RootError error

	Messages []string
}

// Push starts a new stack or appends message to the existing stack
func Push(err error, message string) error {
	var newStack errorStack

	// if the incoming error is a errorStack, append our message
	if stack, ok := err.(errorStack); ok {
		newStack.RootError = stack.RootError
		newStack.Messages = append(stack.Messages, message)
	} else {
		newStack.RootError = err
		newStack.Messages = []string{message}
	}

	return newStack
}

// Pushf starts a new stack or appends a formatted message to the existing stack
func Pushf(err error, format string, params ...[]string) error {
	message := fmt.Sprintf(format, params)
	return Push(err, message)
}

// Error implements the 'error' interface by Joining the stack using the
// default separator
func (stack errorStack) Error() string {
	return stack.Join(defaultSeparator)
}

// Root returns the original error
func (stack errorStack) Root() error {
	return stack.RootError
}

// Stack returns a slice of strings that are in the order of the call stack.
// The final string is the output of Error() from the root error
func (stack errorStack) Stack() []string {
	messagesSize := len(stack.Messages)
	result := make([]string, messagesSize+1)

	// put the root message at the end
	result[messagesSize] = stack.RootError.Error()

	// store the other messages in the order they came in
	for i := 0; i < messagesSize; i++ {
		result[messagesSize-i-1] = stack.Messages[i]
	}

	return result
}

// Join returns a string created by by joining the output of Stack using the
// separator
func (stack errorStack) Join(sep string) string {
	return strings.Join(stack.Stack(), sep)
}
