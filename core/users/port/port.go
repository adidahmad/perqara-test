package port

import (
	"github.com/adidahmad/perqara-test/core/users/domain"
	"github.com/adidahmad/perqara-test/core/users/entity"
)

type IUsersService interface {
	GetList() ([]domain.Users, error)
	GetById(id string) (domain.Users, error)
	Create(data domain.CreateUserRequest) (domain.Users, error)
	Update(id string, data domain.UpdateUserRequest) (domain.Users, error)
	DeleteById(id string) error
}

type IUsersRepository interface {
	FindAll() ([]*entity.Users, error)
	FindById(id string) (entity.Users, error)
	Insert(data entity.Users) (entity.Users, error)
	Update(id string, data entity.Users) (entity.Users, error)
	DeleteById(id string) error
}
