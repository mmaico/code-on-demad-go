package view

import (
	"code-on-demand-go/contacts/domain"
	"code-on-demand-go/contacts/view/support"
	"github.com/gin-gonic/gin"
	"github.com/nvellon/hal"
	"strconv"
)

type ContactHandler interface {
}

type contactHandler struct {
	contact domain.Contact
}

type ContactResource struct {
	Id     int32         `json:"id"`
	Mobile string        `json:"mobile"`
	Status domain.Status `json:"status"`
}

func (resource ContactResource) newContactResource(contact *domain.Contact) ContactResource {
	return ContactResource{
		Id:     contact.Id,
		Mobile: support.MaskData(contact.Mobile),
		Status: contact.Status,
	}
}

func NewContactHandler(contact domain.Contact) *contactHandler {
	return &contactHandler{
		contact: contact,
	}
}

func (contact contactHandler) FindAll(c *gin.Context) {
	contacts := contact.contact.FindAll().OnlyActive()
	var resources = make([]*hal.Resource, len(contacts))

	for i := 0; i < len(contacts); i++ {
		resources[i] = hal.NewResource(ContactResource{}.newContactResource(contacts[i]),
			"/customers/contacts/"+strconv.Itoa(int(contacts[i].Id)))
	}
	c.JSON(200, resources)
}
