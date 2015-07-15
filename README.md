## Functional Programming Constructs in Golang

These functions where coded to understand the difficulty of implementing these functions in golang. 
I do not recommend you use it. A simple for loop will do and work in most of the cases.

#### Check out the test file to see the usage of these

### FUNCTIONS

```go

// Definition of Filter function
type filterFunc func(interface{}) bool

func Filter(fn filterFunc, array interface{}) []interface{}
    Filter filters the array based on the predicate
```

```go

// Definition of Foldl function
type foldlFunc func(interface{}, interface{}) interface{}

func Foldl(fn foldrFunc, array interface{}, accumulator interface{}) interface{}
    Folds left the array values (reduction) based on the function
```

```go

// Definition of Map function
type mapFunc func(interface{}) interface{}

func Map(fn mapFunc, array interface{}) []interface{}
    Map maps the function onto the array
```


