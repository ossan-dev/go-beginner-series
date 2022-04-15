## requirements
1. VS Code
1. Go installed
1. staticcheck
1. go fmt
1. Postman (or any other software for testing API calls)

## setup
1. run `cd src/`
1. run `go mod init todoapi` and it will return =>  go: creating new go.mod: module todoapi
1. confirm that the file will look like this:
```go
module todoapi

go 1.18

```
1. create "main.go"
1. fill in with this code  
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```
1. issue `go mod tidy`
1. issue `go run .` or press F5 (to debug)

## http & routing
### router
1. create "/router" folder
1. create "router.go" file
1. fill in the file with this code:
```go
package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	}).Methods(http.MethodGet)
}
```
1. issue `go get github.com/gorilla/mux` => go: added github.com/gorilla/mux v1.8.0
1. check go.mod to see if the dependency has been added

## formatting
1. `staticcheck ./...`
1. `go fmt ./...` => main.go, router.go