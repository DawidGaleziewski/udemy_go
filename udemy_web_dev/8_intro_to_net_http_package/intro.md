most important thing to understand is the Handler interface

!IMPORTANT
```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```
as in any interface, implementing this method will bring with it the Handler interface. Therefore it will implicitly have handler interface.


second most important thing is the http.ListenAndServe function from this package
```go
var d somethingThatIsOfHenderInterface
http.ListenAndServe(":8080", d)
```