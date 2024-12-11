package repository

import (
	"a21hc3NpZ25tZW50/model"
	"fmt"
	"gorm.io/gorm"
)

type UserRepository interface {
	Add(user model.User) error
	CheckAvail(user model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepository {
	return &userRepository{db}
}
func (u *userRepository) Add(user model.User) error {
	err := u.db.Create(&user).Error
	return err
}

func (u *userRepository) CheckAvail(user model.User) error {
    if user.Username == "" || user.Password == "" {
        return fmt.Errorf("username or password cannot be empty")
    }
    var existingUser model.User
    err := u.db.Where("username = ? AND password = ?", user.Username, user.Password).First(&existingUser).Error
    if err != nil {
        return err
    }
    return nil
}

