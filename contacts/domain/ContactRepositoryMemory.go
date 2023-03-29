package domain

type ContactRepository interface {
	FindAll() []*Contact
}

type contactRepository struct{}

func NewContactRepository() ContactRepository {
	return &contactRepository{}
}

func (impl contactRepository) FindAll() []*Contact {
	return []*Contact{{Id: 1, Mobile: "11987639875", Status: ACTIVE},
		{2, "1156194823", ACTIVE},
		{3, "1188474748", INACTIVE},
		{4, "1194894847", ACTIVE},
	}
}
