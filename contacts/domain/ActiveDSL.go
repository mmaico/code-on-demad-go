package domain

type OnlyActive struct {
	repository ContactRepository
}

func (contact Contact) FindAll() OnlyActive {
	return OnlyActive{
		repository: contact.Repository,
	}
}
