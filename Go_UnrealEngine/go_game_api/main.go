package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	port_http = flag.Int("port_http", 9088, "The http server port")
)

func main() {
	RunHttpServer()
}

func RunHttpServer() {
	r := gin.Default()
	r.GET("/", HttpAPI_StartRun)
	r.Run(fmt.Sprintf(":%d", *port_http))
}

func HttpAPI_StartRun(c *gin.Context) {
	txt := "Run success"
	c.String(http.StatusOK, txt)
}

/*
> go run main.go
> go build
*/
