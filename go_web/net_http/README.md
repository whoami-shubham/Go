``` Go
  package "net/http"
```
## Handler

It handle the requests

[http.Handler](https://godoc.org/net/http#Handler)
``` Go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```
Any type that implements above ServeHTTP(ResponseWriter,*Request) method is also a handler

## Server

[http.ListenAndServe](https://godoc.org/net/http#ListenAndServe)
``` Go
func ListenAndServe(addr string, handler Handler) error
```
