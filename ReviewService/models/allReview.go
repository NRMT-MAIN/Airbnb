package models

type SingleReview struct {
	ID         int     
	BookingID  int    
	Comment    string 
	Rating     int    
	Updated_At string
	Created_At string
}