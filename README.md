## Functional Programming Constructs in Golang

### FUNCTIONS

```go
func Filter(array interface{}, fn filterFunc) []interface{}
    Filter filters the array based on the predicate
```

```go
func Foldl(array interface{}, fn foldrFunc, accumulator interface{}) interface{}
    Folds left the array values (reduction) based on the function
```

```go
func Map(array interface{}, fn mapFunc) []interface{}
    Map maps the function onto the array
```


