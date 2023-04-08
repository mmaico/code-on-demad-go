package infrastructure

import "code-on-demand-go/contacts/domain"

type contactRepository struct{}

func NewContactRepository() contactRepository {
	return contactRepository{}
}

func (impl contactRepository) FindAll() []*domain.Contact {
	return []*domain.Contact{
		{1, "11987639875", domain.ACTIVE, impl},
		{2, "11561948230", domain.ACTIVE, impl},
		{3, "11884747489", domain.INACTIVE, impl},
		{4, "11948948471", domain.ACTIVE, impl},
	}
}
