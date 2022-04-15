# Go API tutorial: the begin

## Hashtags

golang, http, API, vscode, go, webdev

Welcome back guys 👋 It's been a while since my last post written for our **amazing** community. A lot of cool things have happened during this time (including my transfer to another company). Due to this now, my daily work is focused on Go programming language. I want to say that I'm not yet a well-rounded programmer with this language but I'll do my best to build a very interesting series about the thing that I love most: **write secure, testable and successful software**.  
I'm really excited for this series and I hope it's the same for you, dear reader 😃  

## Preamble
As I said above I'm not the most skilled person at Golang on the Earth, so if u find any mistakes or something that can be improved or explained better let me know and I'll take the necessary actions to fix it.  
Another thing that I want to state is that I strive to have an incremental approach throughout this tutorial, so every post will do only one thing (like the Single Responsibility Principle which is one of my favorite principle when it comes to programming 🤖).  
This tutorial is not written on stones so, if u'd like to see something feel free to reach me and I can introduce your request in the series.

## The big picture 🌎
The final goal of this tutorial is to build a REST API with Go programming language. This API will be a TODO API (I know, I know... very few imagination 🙄). But during our journey we're going to introduce a lot of cool stuff and obviously everything will be backed up from Unit Tests which is an evergreen topic to cover IMO.

## First step 🔰
This post will cover just the basics of how to setup a Golang project and some basic HTTP stuff, such as exposing a dummy HTTP endpoint with a simple GET method. This will be the foundation for the next parts 🏰.

## Let's start 🚀
Without having much more delay we can start having fun together!

### Prerequisites
To follow this tutorial u must have these programs and tools on your machine:
1. Visual Studio Code (or any other IDE)
1. Go installed. U can download from [here](https://go.dev/dl/)
1. staticcheck tool. It's a tool for linting which can perform some checks on your code about simplifications, styling rules, performance issues and so on. U can find more about it [here](https://staticcheck.io/). It's not mandatory but I *strongly* recommend it
1. gofmt is a package that can assist u in formatting the code. I suggest you to install globally on your machine from this [link](https://pkg.go.dev/cmd/gofmt). Again, *not mandatory but suggested*!
1. Postman (or any other API client that you're confortable with)

### Setup 🛠️
Let's create a new folder for this tutorial with some meaningful name (such as *todo-api*). Open this folder with your IDE and inside the integrated terminal issue these two commands: `mkdir src` and `cd src/`.  
Next step is to initialize a Golang package. In order to do this we need this command `go mod init todoapi`. Thanks to this now we're also able to manage the module dependencies by importing go packages from the internet 🌏. If everything goes fine you will see a new file called **go.mod** below the "src/" folder. The content should be something similar to this:
```go
module todoapi

go 1.18
```
🔴 **IMPORTANT**: 1.18 means the *miminum* version needed for compile this Golang program.
## Main
The **main** is the entry point of our program and its **main** function will be the first called by the running OS when it will execute our code  
🧐*NOTE*: Golang is a **compiled** programming language that means that it will produce an executable file to run.  
Now you can create the "main.go" under the "/src" folder. Fill in with this simple code to check that everything is fine:
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```
Confirm that everything is working by issuing in the CLI this command `go run .` or by pressing F5 (start debugging) in your IDE. You should see the string `Hello, World!` 👋 printed in the console.
## Routing
As you now, web api is all about receiving an HTTP request and providing and HTTP response. When an HTTP request comes in we need some mechanism to decide which controllers (or handlers, these two names can be used interchangeably but I prefer to stick with the former) have to manage this kind of request and provide a relevant response.  
This mechanism is called **Routing**. This one has to instrumented on how to manage the requests so we need a way of mapping each accepted route with relative controller within a one-to-one relationship 🚧.  
In order to setup the router u need to perform the following steps:
1. Create a "router/" folder below "src/"
1. Create a "router.go" under this folder
1. Fill the router.go file with this code:
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
    This code will instrument the HTTP server to manage the requests made to "/ping" 🏓 endpoint with the GET HTTP method and it provides an inline handler that will only respond with a dummy plain text "pong" 🏓 and 200 Status OK.  
    An **handler** is a function that has the following two parameters in the method signature: ResponseWriter and Request. The former is the actual response that will be sent back to client 🔙 and so you've to manipulate it (body, status code, headers, etc) while the latter is the HTTP request that comes in 🔜 and it holds all of the info such as request payload, route, querystring params and so on which you can access.  
    There are a lot of packages that can be used for routing, my preferred one is the one from Gorilla 🦍. More details [here](https://github.com/gorilla/mux).
1. In the CLI issue the following command to install this library `go get github.com/gorilla/mux`
1. Check if this dependency was added in go.mod file
## Golang tools
#### staticcheck
Now it's time to lint your code 👨‍🏫. Open up your CLI and issue this command `staticcheck ./...`.
#### gofmt
After linting it's time for formatting 💠. In the CLI issue `go fmt ./...` and you will see the files that have been formatted according to Golang formatting rules.  
⚠️ **WARNING**: u need to make sure that you're located where your main.go file is when running these commands otherwise you end up in having a lot of issues due to the fact that the tools either doesn't recognize it as a valid Golang pkg or nothing happened at all.
## Wiring all together ⛓️
Now, switch back to main.go and make use of newly created stuff through this code:
```go
package main

import (
	"fmt"
	"net/http"

	"todoapi/router"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("starting up HTTP server...")

	r := mux.NewRouter()

	router.SetupRoutes(r)

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		panic(err)
	}
}
```
Here we're going to *instantiate a new router* with a method provided by Gorilla package. After that we instrument it with the setup method in our router.go file.  
The last part is responsible for starting up an HTTP server that will accept requests on port 8000 and it will manage them according to **r** that is our router BTW.
## Test 🤓
For this time (and only for this time) we're going to do some manual tests. Don't worry from the next post we're gonna to use unit tests 😉.  
First of all launch our Go package with `go run .` or with F5. Next, open Postman and prepare a simple GET request pointing to this URL: "http://localhost:8000/ping". Issue the request and you should get back a "pong" response with 200 HTTP StatusOK 👌.
👀 **NOTES**: to interrupt the running server you can press CTRL + C on your keyboard.
## Success 🐎
We have successfully finished up this first tutorial which is just a **warm-up** for the series. Now you have the basics about how to spin up an HTTP server in Golang and instrument it to respond to the requests.
## What's next ⏭️
The next article will be about our controller 🛂. So we're going to define our first controller that will interact with HTTP request and HTTP response and we will introduce the unit test that it's a huge topic to deal with ⛰️.  
If this sounds interesting for you don't miss if for any reason 😁.  

I hope you enjoy this post and find it useful. If you have any questions or you want to spot me some errors I really appreciate it and I'll make my best to follow up. If you enjoy it and would like to sustain me consider giving a like and sharing on your favorite socials. If u want u can add me on your socials this makes me very very happy!

Stay safe and see you soon! 😎