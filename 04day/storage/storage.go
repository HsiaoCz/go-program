package storage

import "go-program/04day/types"

type Storage interface {
	Get(int) *types.User
}
