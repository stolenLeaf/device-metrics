package usertenants

import "gorm.io/gorm"

type UserTenants struct{
	gorm.Model
	Uuid string `gorm:"size:100;not null"`
	UserId string `gorm:"size:100;not null"`
	Role string `gorm:"size:100;not null"`
	Status string `gorm:"size:100;not null"`

}
