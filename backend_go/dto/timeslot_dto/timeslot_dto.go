package timeslot_dto

type CreateTimeslotRequest struct {
	StartTime    string  `json:"start_time" binding:"required"`
	EndTime      string  `json:"end_time" binding:"required"`
	PriceWeekday float64 `json:"price_weekday" binding:"required"`
	PriceWeekend float64 `json:"price_weekend" binding:"required"`
}

type UpdateTimeslotRequest struct {
	StartTime    string  `json:"start_time"`
	EndTime      string  `json:"end_time"`
	PriceWeekday float64 `json:"price_weekday"`
	PriceWeekend float64 `json:"price_weekend"`
}
