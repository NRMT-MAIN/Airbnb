package models

type Review struct {
	ID         int64     
	BookingID  int    
	Comment    string 
	Rating     int    
	Updated_At string
	Created_At string   
}

