package booking_dto

type CreateBookingRequest struct {
	CourtID     uint   `json:"court_id" binding:"required"`
	TimeslotID  uint   `json:"timeslot_id" binding:"required"`
	BookingDate string `json:"booking_date" binding:"required"`
}

type UpdateBookingRequest struct {
	Status       string `json:"status"`
	PaymentToken string `json:"payment_token"`
}
