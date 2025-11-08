package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name  string `gorm:"unique;not null;type:varchar(50)" json:"name"`
	Users []User `gorm:"foreignKey:RoleID" json:"user,omitempty"`
}
