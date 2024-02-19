package service

import (
	"github.com/adidahmad/perqara-test/core/users/domain"
	"github.com/adidahmad/perqara-test/core/users/entity"
	usersPort "github.com/adidahmad/perqara-test/core/users/port"
	"golang.org/x/crypto/bcrypt"
)

type UsersService struct {
	UsersRepository usersPort.IUsersRepository
}

func NewUsersService(usersRepo usersPort.IUsersRepository) usersPort.IUsersService {
	return UsersService{
		UsersRepository: usersRepo,
	}
}

func (s UsersService) GetList() ([]domain.Users, error) {
	res, err := s.UsersRepository.FindAll()
	if err != nil {
		return []domain.Users{}, err
	}

	users := []domain.Users{}
	for _, user := range res {
		users = append(users, domain.Users(*user))
	}

	return users, nil
}

func (s UsersService) GetById(id string) (domain.Users, error) {
	user, err := s.UsersRepository.FindById(id)
	if err != nil {
		return domain.Users{}, err
	}

	return domain.Users(user), nil
}

func (s UsersService) Create(data domain.CreateUserRequest) (domain.Users, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return domain.Users{}, err
	}

	entity := entity.Users{
		Email:    data.Email,
		Password: string(hashedPassword),
		IsActive: data.IsActive,
	}
	res, err := s.UsersRepository.Insert(entity)
	if err != nil {
		return domain.Users{}, err
	}

	return domain.Users(res), nil
}

func (s UsersService) Update(id string, data domain.UpdateUserRequest) (domain.Users, error) {
	entity := entity.Users{
		Email:    data.Email,
		Password: data.Password,
		IsActive: data.IsActive,
	}
	res, err := s.UsersRepository.Update(id, entity)
	if err != nil {
		return domain.Users{}, err
	}

	return domain.Users(res), nil
}

func (s UsersService) DeleteById(id string) error {
	return s.UsersRepository.DeleteById(id)
}
