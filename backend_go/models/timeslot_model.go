package models

import "gorm.io/gorm"

type Timeslot struct {
	gorm.Model
	StartTime    string    `gorm:"not null;type:time" json:"start_time"`
	EndTime      string    `gorm:"not null;type:time" json:"end_time"`
	PriceWeekday float64   `gorm:"not null" json:"price_weekday"`
	PriceWeekend float64   `gorm:"not null" json:"price_weekend"`
	Bookings     []Booking `gorm:"foreignKey:TimeslotID" json:"booking,omitempty"`
}
