package main

import (
	"code-on-demand-go/contacts/view"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	view.NewContactRoutes(server).ContactRoutes()
	server.Run(":8181")
}
