package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")

	userR := v1.Group("/user")
	userR.POST("/register")

	v1Auth := v1.Group("/auth")

	v1Auth.Use()
	v1Auth.POST()
}
