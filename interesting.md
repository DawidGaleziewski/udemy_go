(EBNF) Extended Backus–Naur form. Describes syntax
https://en.wikipedia.org/wiki/Extended_Backus%E2%80%93Naur_form

golang.org and golang.doc are important sources of information


Advanced go courses:
https://www.ardanlabs.com/


## coma ok idiom
coma ok idiom (, ok) is a popular idiomatic structure used to check the value before doing something with it
```go
    if v, ok := mapValue["key"]; ok {
        //do something
    }
```


## aliasing a type

Aliasing a type with a primitive type is considered a bad practice most of the time
```go
type foo int
var y foo
```