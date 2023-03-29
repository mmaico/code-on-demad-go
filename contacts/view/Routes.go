package view

import (
	"code-on-demand-go/contacts/application"
	"code-on-demand-go/contacts/domain"
	"github.com/gin-gonic/gin"
)

type contactRoute struct {
	handler contactHandler
	engine  *gin.Engine
}

func NewContactRoutes(engineParam *gin.Engine) *contactRoute {
	return &contactRoute{
		handler: *NewContactHandler(application.ContactFacade(domain.NewContactRepository())),
		engine:  engineParam,
	}
}

func (route contactRoute) ContactRoutes() {
	route.engine.GET("/customers/contacts", route.handler.FindAll)
}
