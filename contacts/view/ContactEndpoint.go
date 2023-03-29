package view

import (
	"code-on-demand-go/contacts/domain"
	"github.com/gin-gonic/gin"
)

type ContactHandler interface {
}

type contactHandler struct {
	contact domain.Contact
}

func NewContactHandler(contact domain.Contact) *contactHandler {
	return &contactHandler{
		contact: contact,
	}
}

func (contact contactHandler) FindAll(c *gin.Context) {
	contacts := contact.contact.FindAll().OnlyActive()
	c.JSON(200, contacts)
}
