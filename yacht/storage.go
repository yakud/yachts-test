package yacht

import "fmt"

type Storage struct {
}

func (t *Storage) Add(yacht *Model) error {
	fmt.Printf("%+v\n", yacht)
	return nil
}

func NewStorage() *Storage {
	return &Storage{}
}
