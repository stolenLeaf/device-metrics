package usertenants

import "github.com/stolenleaf/device-metrics/internal/database"

type UserTenantsRepo struct{}

func (r *UserTenantsRepo) GetAll() ([]UserTenants, error) {
	var usertenants []UserTenants
	err := database.DB.Find(&usertenants).Error

	return usertenants, err
}

func (r *UserTenantsRepo) Create(u *UserTenants) error {
	return database.DB.Create(u).Error
}
