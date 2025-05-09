package repository

import "github.com/w-tyler/mcp-test/internal/domain/entity"

type UserRepository interface {
	GetByID(id string) (*entity.User, error)
	List() ([]*entity.User, error)
}
