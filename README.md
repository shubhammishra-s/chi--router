# chi - web router
A lightweight, idiomatic and composable ```router``` for building Go HTTP services.

# To use

make sure you have intialized your root module like ```github.com/shubhammishra-1```

install ```chi``` 

```go get -u github.com/go-chi/chi/v5```


# chi.NewRouter() *Mux

NewRouter() returns a new ```Mux``` object that implements the Router interface. 

Mux is a simple HTTP route multiplexer that parses a request path, records any URL params, and executes an end handler. It implements the http.Handler interface and is friendly with the standard library.
Mux is designed to be fast, minimal and offer a powerful API for building modular and composable HTTP services with a large set of handlers. It's particularly useful for writing large REST API services that break a handler into many smaller parts composed of middlewares and end handlers.

```go 
router := chi.NewRouter()
```

# Route handling 
Route handling means how an application's endpoints (URIs) respond to client requests.
`chi` allows you to route/handle any HTTP request method, such as all the usual suspects: `GET`, `POST`, `HEAD`, `PUT`, `PATCH`, `DELETE`, `OPTIONS`, `TRACE`, `CONNECT`...

## router.Get(pattern string, handlerFn http.HandlerFunc)

Get() adds the route `pattern` that matches a GET http method to execute the `handlerFn` http.HandlerFunc.

## router.Delete(pattern string, handlerFn http.HandlerFunc)

Delete() adds the route `pattern` that matches a DELETE http method to execute the `handlerFn` http.HandlerFunc. 

## router.Patch(pattern string, handlerFn http.HandlerFunc)

Patch adds the route `pattern` that matches a PATCH http method to execute the `handlerFn` http.HandlerFunc. 

## router.Post(pattern string, handlerFn http.HandlerFunc)

Post() adds the route `pattern` that matches a POST http method to execute the `handlerFn` http.HandlerFunc. 

## router.Put(pattern string, handlerFn http.HandlerFunc)

Put adds the route `pattern` that matches a PUT http method to execute the `handlerFn` http.HandlerFunc.

## router.Connect(pattern string, handlerFn http.HandlerFunc)

Connect() adds the route `pattern` that matches a CONNECT http method to execute the `handlerFn` http.HandlerFunc. 

# About http.HandlerFunc
The HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers. If f is a function with the appropriate signature, HandlerFunc(f) is a Handler that calls f.

for example :::
```go 
func(w http.ResponseWriter, r *http.Request) {
	...
}
```

# About Middlewares
Middleware performs some specific function on the HTTP request or response at a specific stage in the HTTP pipeline before or after the user defined controller. Middleware is a design pattern to eloquently add cross cutting concerns like logging, handling authentication without having many code contact points.

in ```chi``` package there are some middlewares which can be used..

to use middlewares make sure you have imported ```"github.com/go-chi/chi/v5/middleware"``` in your go program.

will discuss some very common used middlewares line by line...

```middleware.Logger``` 
Logger is a middleware that logs the start and end of each request, along with some useful data about what was requested, what the response status was, and how long it took to return. When standard output is a TTY, Logger will print in color, otherwise it will print in black and white. Logger prints a request ID if one is provided.

```middleware.RequestID```
RequestID is a middleware that injects a request ID into the context of each request. A request ID is a string of the form "host.example.com/random-0001", where "random" is a base62 random string that uniquely identifies this go process, and where the last number is an atomically incremented request counter.

```middleware.Recoverer```
Recoverer is a middleware that recovers from panics, logs the panic (and a backtrace), and returns a HTTP 500 (Internal Server Error) status if possible. Recoverer prints a request ID if one is provided.

```middleware.RedirectSlashes```
RedirectSlashes is a middleware that will match request paths with a trailing slash and redirect to the same path, less the trailing slash. 

```middleware.Timeout(t)```
Timeout is a middleware that cancels ctx after a given timeout and return a 504 Gateway Timeout error to the client.
It's required that you select the ctx.Done() channel to check for the signal if the context has reached its deadline and return, otherwise the timeout signal will be just ignored.

# router.Use(middlewares ...func(http.Handler) http.Handler)

Use() appends a middleware handler to the Mux middleware stack.
The middleware stack for any Mux will execute before searching for a matching route to a specific handler, which provides opportunity to respond early, change the course of the request execution, or set request-scoped values for the next http.Handler.

for example::: To use any middleware  ```router.Use(middleware.Logger)```


# Sub routing 
there are various way for sub routing... one of using Mount function

````router.Mount(pattern string, handler http.Handler)```` 
Mount attaches another http.Handler or chi Router as a subrouter along a routing path. It's very useful to split up a large API as many independent routers and compose them as a single service using Mount. See _examples/.
Note that Mount() simply sets a wildcard along the `pattern` that will continue routing at the `handler`, which in most cases is another chi.Router. As a result, if you define two Mount() routes on the exact same pattern the mount will panic.

# References

```https://go-chi.io```

```https://pkg.go.dev/github.com/go-chi/chi/v5```

```https://pkg.go.dev/github.com/go-chi/chi/v5/middleware```
