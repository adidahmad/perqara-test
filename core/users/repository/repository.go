package repository

import (
	"errors"

	"github.com/adidahmad/perqara-test/core/users/entity"
	usersPort "github.com/adidahmad/perqara-test/core/users/port"

	"gorm.io/gorm"
)

type UsersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) usersPort.IUsersRepository {
	return UsersRepository{
		db: db,
	}
}

func (r UsersRepository) FindAll() ([]*entity.Users, error) {
	users := []*entity.Users{}

	err := r.db.Find(&users).Error
	if err != nil {
		return []*entity.Users{}, err
	}

	return users, nil
}

func (r UsersRepository) FindById(id string) (entity.Users, error) {
	user := entity.Users{}

	err := r.db.First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.Users{}, err
	}

	return user, nil
}

func (r UsersRepository) FindByEmail(email string) (entity.Users, error) {
	user := entity.Users{}

	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return entity.Users{}, err
	}

	return user, nil
}

func (r UsersRepository) Insert(data entity.Users) (entity.Users, error) {
	err := r.db.Create(&data).Error
	if err != nil {
		return entity.Users{}, err
	}

	return data, nil
}

func (r UsersRepository) DeleteById(id string) error {
	user, err := r.FindById(id)
	if err != nil {
		return err
	}

	return r.db.Delete(&user).Error
}

func (r UsersRepository) DeleteByIds(ids []string) error {
	user := []entity.Users{}
	err := r.db.Find(&user, ids).Error
	if err != nil {
		return err
	}

	return r.db.Delete(&user).Error
}

func (r UsersRepository) Update(id string, data entity.Users) (entity.Users, error) {
	user := entity.Users{}
	err := r.db.Find(&user, id).Error
	if err != nil {
		return entity.Users{}, err
	}

	err = r.db.Model(&user).Updates(&data).Error
	if err != nil {
		return entity.Users{}, err
	}

	return user, nil
}
