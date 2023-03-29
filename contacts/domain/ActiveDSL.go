package domain

type OnlyActive struct{}

func (contact Contact) FindAll() OnlyActive {
	return OnlyActive{}
}
