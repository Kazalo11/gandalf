package server

import "github.com/gin-gonic/gin"

func Start() {
	router := gin.Default()
	router.Run()
}