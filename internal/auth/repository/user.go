package repository

import (
	"github.com/google/uuid"
	"github.com/your-moon/grape-shared/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByID(id uuid.UUID) (*entities.UserUUID, error) {
	var user entities.UserUUID
	err := r.db.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) List() ([]*entities.UserUUID, error) {
	var users []*entities.UserUUID
	err := r.db.Order("created_at desc").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) FindByPhone(phone string) (*entities.UserUUID, error) {
	var user entities.UserUUID
	err := r.db.Where("phone_number = ?", phone).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByEmail(email string) (*entities.UserUUID, error) {
	var user entities.UserUUID
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByPhoneOrEmail(identifier string) (*entities.UserUUID, error) {
	var user entities.UserUUID
	err := r.db.Where("phone_number = ? OR email = ?", identifier, identifier).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Create(user *entities.UserUUID) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) Update(user *entities.UserUUID) error {
	return r.db.Model(&entities.UserUUID{}).Where("id = ?", user.ID).Updates(user).Error
}

func (r *UserRepository) Delete(user *entities.UserUUID) error {
	return r.db.Delete(user).Error
}
