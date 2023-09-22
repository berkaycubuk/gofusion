package gofusion

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
)

type Reactor struct {
	Router *gin.Engine
	Scheduler *gocron.Scheduler
}

func Init() *Reactor {
	router := gin.Default()
	scheduler := gocron.NewScheduler(time.UTC)

	return &Reactor{
		Router: router,
		Scheduler: scheduler,
	}
}

func Run(reactor *Reactor, port int) {
	reactor.Scheduler.StartAsync()
	http.ListenAndServe(fmt.Sprintf(":%v", port), reactor.Router)
}
