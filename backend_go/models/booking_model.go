package models

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	UserID       uint      `gorm:"not null" json:"-"`
	User         User      `gorm:"foreignKey:UserID" json:"user"`
	CourtID      uint      `gorm:"not null" json:"-"`
	Court        Court     `gorm:"foreignKey:CourtID" json:"court"`
	TimeslotID   uint      `gorm:"not null" json:"-"`
	Timeslot     Timeslot  `gorm:"foreignKey:TimeslotID" json:"timeslot"`
	BookingDate  time.Time `gorm:"not null;type:date" json:"booking_date"`
	Status       string    `gorm:"not null;type:varchar(20);default:'pending'" json:"status"`
	TotalPrice   float64   `gorm:"not null" json:"total_price"`
	PaymentToken string    `gorm:"type:text" json:"payment_token,omitempty"`
}
