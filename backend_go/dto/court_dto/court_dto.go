package court_dto

type CreateCourtRequest struct {
	CourtName string `json:"court_name" binding:"required"`
	Type      string `json:"type" binding:"required"`
	Location  string `json:"location" binding:"required"`
}

type UpdateCourtRequest struct {
	CourtName string `json:"court_name"`
	Type      string `json:"type"`
	Location  string `json:"location"`
}
