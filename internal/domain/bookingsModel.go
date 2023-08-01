package domain

import (
	"github.com/gofrs/uuid"
)



type Booking struct {
	Id 					uuid.UUID 
	ClientName			string		`json:"clientName"`
	PhotographerName	string		`json:"photographerName"`
	Package 			string		`json:"package"`
	DateTime 			string 		`json:"dateTime"`
	Location 			string		`json:"location"`
	Status				string		`json:"status"`

}

type BookingRequestFormat struct {
	Data         []Booking	 `json:"data"`
}