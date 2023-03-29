package infrastructure

import "code-on-demand-go/contacts/domain"

type contactRepository struct{}

func NewContactRepository() contactRepository {
	return contactRepository{}
}

func (impl contactRepository) FindAll() []*domain.Contact {
	return []*domain.Contact{
		{1, "11987639875", domain.ACTIVE, impl},
		{2, "1156194823", domain.ACTIVE, impl},
		{3, "1188474748", domain.INACTIVE, impl},
		{4, "1194894847", domain.ACTIVE, impl},
	}
}
