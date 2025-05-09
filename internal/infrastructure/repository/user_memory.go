package repository

import (
	"errors"
	"github.com/w-tyler/mcp-test/internal/domain/entity"
)

type UserMemoryRepo struct {
	users map[string]*entity.User
}

func NewUserMemoryRepo() *UserMemoryRepo {
	return &UserMemoryRepo{users: make(map[string]*entity.User)}
}

func (r *UserMemoryRepo) GetByID(id string) (*entity.User, error) {
	u, ok := r.users[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return u, nil
}

func (r *UserMemoryRepo) List() ([]*entity.User, error) {
	var list []*entity.User
	for _, u := range r.users {
		list = append(list, u)
	}
	return list, nil
}
