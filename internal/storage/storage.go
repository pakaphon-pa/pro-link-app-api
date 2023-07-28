package storage

type (
	Storage struct {
	}
)

func (s *Storage) NewStorage() *Storage {
	return &Storage{}
}
