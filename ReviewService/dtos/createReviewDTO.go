package dtos

type CreateReviewDTO struct {
	BookingID   string  `json:"bookingId"`
	Comment     string  `json:"comment"`
	Rating      int   `json:"rating"`
}