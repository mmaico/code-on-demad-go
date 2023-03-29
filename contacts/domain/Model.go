package domain

type Contact struct {
	Id         int32
	Mobile     string
	Status     Status
	Repository ContactRepository
}

type Status int8

const (
	INACTIVE Status = 0
	ACTIVE   Status = 1
)

func NewContact(repository ContactRepository) Contact {
	return Contact{
		Repository: repository,
	}
}
func (onlyActive OnlyActive) OnlyActive() []*Contact {
	return onlyActive.repository.FindAll()
}

type ContactRepository interface {
	FindAll() []*Contact
}
