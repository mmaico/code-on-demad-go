package view

import (
	"code-on-demand-go/contacts/domain"
	"code-on-demand-go/contacts/infrastructure"
	"github.com/gin-gonic/gin"
)

type contactRoute struct {
	handler contactHandler
	engine  *gin.Engine
}

func NewContactRoutes(engineParam *gin.Engine) *contactRoute {

	return &contactRoute{
		handler: *NewContactHandler(domain.NewContact(infrastructure.NewContactRepository())),
		engine:  engineParam,
	}
}

func (route contactRoute) ContactRoutes() {
	route.engine.GET("/customers/contacts", route.handler.FindAll)
}
