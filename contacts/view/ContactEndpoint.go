package view

import (
	"code-on-demand-go/contacts/application"
	"github.com/gin-gonic/gin"
)

type ContactHandler interface {
}

type contactHandler struct {
	service application.ContactFacade
}

func NewContactHandler(service application.ContactFacade) *contactHandler {
	return &contactHandler{
		service: service,
	}
}

func (contactHandler *contactHandler) FindAll(c *gin.Context) {
	contacts := contactHandler.service.FindAll()
	c.JSON(200, contacts)
}
