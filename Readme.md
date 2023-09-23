# GoFusion

GoFusion is a simple approach to build web services. Right now, it's just a glue between gin and gorm.

> Disclaimer: This project is under development and should not be used in production environments, so use it with caution!

## How to use it?

First, get the package to your project:
```bash
go get github.com/berkaycubuk/gofusion
```

Simple example:
```go
package main

func main() {
    // initialize gofusion
    reactor := gofusion.Init()

    // you can access the router with reactor.Router and it's a *gin.Engine, so you can use it like you do with go-gin

    // run the http server at port 4000
    gofusion.Run(reactor, 4000)
}
```

You may want to use CORS middleware coming with gofusion:
```go
reactor.Router.Use(gofusion.CORSMiddleware())
```

GoFusion also comes with scheduling thanks to gocron. It is stored in Reactor as Scheduler:
```go
reactor.Scheduler.Every(10).Minutes().Do(func() {
    fmt.Println("Hello!")
})
```
