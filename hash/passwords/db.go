package passwords

import (
	"errors"
)

type Database interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
	Delete(key string) error
}

type inMemoryDatabase struct {
	data map[string][]byte
}

func NewInMemoryDatabase() Database {
	return &inMemoryDatabase{data: make(map[string][]byte)}
}

func (s *inMemoryDatabase) Get(key string) ([]byte, error) {
	if value, ok := s.data[key]; ok {
		return value, nil
	}

	return nil, errors.New("key not found")
}

func (s *inMemoryDatabase) Set(key string, value []byte) error {
	s.data[key] = value
	return nil
}

func (s *inMemoryDatabase) Delete(key string) error {

	if _, ok := s.data[key]; ok {
		delete(s.data, key)
		return nil
	}

	return errors.New("key not found")
}
