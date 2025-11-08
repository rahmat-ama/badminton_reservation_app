package models

import "gorm.io/gorm"

type Court struct {
	gorm.Model
	CourtName string    `gorm:"unique;not null;type:varchar(100)" json:"court_name"`
	Type      string    `gorm:"type:varchar(50)" json:"type"`
	Location  string    `gorm:"not null;type:varchar(255)" json:"location"`
	Bookings  []Booking `gorm:"foreignKey:CourtID" json:"booking,omitempty"`
}
