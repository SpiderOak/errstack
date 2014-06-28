errstack
========

a cumulative container for go language errors

A common idiom in go is for a function to return 'error or 'nil',
where 'error' is an object that implements the error interface.

This return value may be passed up a chain of calling functions until
some caller wants to deal with it.

see http://blog.golang.org/error-handling-and-go

It is often useful for intervening functions to add information to the
error before passing it on.

for example:

```go
func UseObject() {
	if object, err := FetchObject(); err != nil  {
		log.Fatalf("UseObject could not FetchObject: %s", err)
	}
}

func FetchOject() (Object, error) {
	if err := CreateObject(); err != nil  {
		nil, fmt.Errorf("FetchObject count not CreateObject: %s", err)
	}
}

func CreatObject() (Object, error) {
	file, err := os.Open("invalid path"); err != nil {
		return nil, err
	}
}
```

This idiom suffers from two drawbacks: 
    1) It's hard to maintain a consistent format 
    2) therre are cases where a calling function wants to examine the original 
       error.

That's is what errstack is for:

```go
func UseObject() {
	if object, err := FetchObject(); err != nil  {
		if pathError, ok := err.Root().(*os.PathError); ok {
			// do something special for PathError
		}
		log.Fatalf("UseObject could not FetchObject: %s", err)
	}
}

func FetchOject() (Object, error) {
	if err := CreateObject(); err != nil  {
		return nil, errstack.Push(err, "FetchObject")
	}
}

func CreatObject() (Object, error) {
	if file, err := os.Open("invalid path"); err != nil {
		return nil, errstack.Push(err, "CreateObject")
	}
}
```
