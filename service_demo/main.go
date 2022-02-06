package main

import (
	"service_demo/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	middleware.ContextGinToController(r).Run()
}
