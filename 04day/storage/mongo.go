package storage

import "go-program/04day/types"

type MongoStorage struct {
}

func (m *MongoStorage) Get(id int) *types.User {
	return &types.User{
		ID:   1,
		Name: "Bob",
	}
}
