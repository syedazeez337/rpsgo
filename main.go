package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	fmt.Println("Server is starting...")

	r := gin.Default()

	r.GET("/hello", hello)
	r.GET("/ping", pong)

	r.Run()

}

type HeaderGin gin.H

var helloMsg = HeaderGin{
	"message": "Hello world",
}

var pongMsg = HeaderGin{
	"message": "pong",
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, helloMsg)
}

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, pongMsg)
}
