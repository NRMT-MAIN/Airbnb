package db

type Storage struct {
	ReviewRepository ReviewRepositorty
}

func NewStorage() *Storage {
	return &Storage{
		ReviewRepository: &ReviewRepositortyImpl{},
	}
}