package storage

import "go-program/04day/types"

type MemoryStorage struct{}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

func (m *MemoryStorage) Get(id int) *types.User {
	return &types.User{
		ID:   1,
		Name: "lisi",
	}
}
