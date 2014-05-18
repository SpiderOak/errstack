package errstack

import (
	"fmt"
	"os"
	"testing"
)

func TestSingleError(t *testing.T) {
	rootMessage := "root"
	err := fmt.Errorf(rootMessage)
	newErr := Push(err, "")

	testString := newErr.Error()
	expectedString := "|root"
	if testString != expectedString {
		t.Fatalf("expecting '%s' found '%s'", expectedString, testString)
	}
}

func TestTwoCalls(t *testing.T) {
	rootMessage := "root"
	rootErr := fmt.Errorf(rootMessage)

	firstMessage := "first"
	secondMessage := "second"

	firstErr := Push(rootErr, firstMessage)
	secondErr := Push(firstErr, secondMessage)

	testString := secondErr.Error()
	expectedString := "second|first|root"
	if testString != expectedString {
		t.Fatalf("expecting '%s' found '%s'", expectedString, testString)
	}
}

func TestRetrieveRoot(t *testing.T) {
	_, rootErr := os.Open("this is not a path")

	firstMessage := "first"
	secondMessage := "second"

	firstErr := Push(rootErr, firstMessage)
	secondErr := Push(firstErr, secondMessage)

	secondErrStack, ok := secondErr.(ErrStack)
	if !ok {
		t.Fatalf("unable to cast secondErr %T", secondErr)
	}

	_, ok = secondErrStack.Root().(*os.PathError)
	if !ok {
		t.Fatalf("unable to retrieve root %T", secondErrStack.Root())
	}
}
