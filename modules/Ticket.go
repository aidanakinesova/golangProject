package modules

import (
	"time"
)

type Ticket struct{
	Id 				int 		`json:"id"`
	FromWhere 		string		`json:"fromWhere" binding:"required"`
	ToWhere 		string  	`json:"toWhere" binding:"required"`
	DepartureDate 	string 		`json:"departureDate" binding:"required"`
	DepartureTime 	time.Time   `json:"departureTime" binding:"required"`
	ArrivalTime 	time.Time   `json:"arrivalTime" binding:"required"`
	Duration 		string  	`json:"duration" binding:"required"`
	Price 			int 		`json:"price" binding:"required"`
}