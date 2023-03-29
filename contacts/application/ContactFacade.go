package application

import (
	"code-on-demand-go/contacts/domain"
)

type ContactFacade interface {
	FindAll() []*domain.Contact
}

type contactModel struct {
	repository domain.ContactRepository
}

func NewContactApplication(repo domain.ContactRepository) ContactFacade {
	return &contactModel{
		repository: repo,
	}
}

func (repo *contactModel) FindAll() []*domain.Contact {
	return repo.repository.FindAll()
}
