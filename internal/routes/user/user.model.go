package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"size:100;not null"`
	LastName string `gorm:"size:100;not null"`
	Office string `gorm:"size:100;not null"`
	LastConnection string `gorm:"size:100;not null"`
	Uuid string `gorm:"size:100;not null"`
}
