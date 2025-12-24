package user

import "github.com/stolenleaf/device-metrics/internal/database"

type UserRepo struct{}

func (r *UserRepo) GetAll() ([]User, error) {
	var users []User
	err := database.DB.Find(&users).Error

	return users, err
}

func (r *UserRepo) Create(u *User) error {
	return database.DB.Create(u).Error
}
