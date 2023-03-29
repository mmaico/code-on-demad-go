package domain

type Contact struct {
	Id     int32
	Mobile string
	Status Status
}

type Status int8

const (
	INACTIVE Status = 0
	ACTIVE   Status = 1
)

func (onlyActive OnlyActive) OnlyActive() []*Contact {
	repository := NewContactRepository()
	return repository.FindAll()
}
