package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string    `gorm:"unique;not null;type:varchar(50)" json:"username"`
	Email    string    `gorm:"not null;type:varchar(100)" json:"email"`
	Password string    `gorm:"not null" json:"-"`
	Kontak   string    `gorm:"type:varchar(20)" json:"kontak"`
	RoleID   uint      `gorm:"not null" json:"-"`
	Role     Role      `gorm:"foreignKey:RoleID" json:"-"`
	Bookings []Booking `gorm:"foreignKey:UserID" json:"booking,omitempty"`
}
