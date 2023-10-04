package repositories

import (
	"myapp/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.User, error)
	GetUser(ID int) (models.User, error)
	AddUser(user *models.User) error
	DeleteUser(ID int) error
	UpdateUser(ID int, user *models.User) error
}

type repository struct {
	db *gorm.DB
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Raw("SELECT * FROM users").Scan(&users).Error

	return users, err
}

func (r *repository) GetUser(ID int) (models.User, error) {
	var user models.User
	err := r.db.Raw("SELECT * FROM users WHERE id=?", ID).Scan(&user).Error

	return user, err
}

func (r *repository) AddUser(user *models.User) error {
	result := r.db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repository) DeleteUser(ID int) error {
	result := r.db.Delete(&models.User{}, ID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repository) UpdateUser(ID int, user *models.User) error {
	result := r.db.Model(&models.User{}).Where("id = ?", ID).Updates(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
