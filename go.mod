module hotel-booking.com

go 1.16

replace hotel-booking.com/v1/app => ./app

require (
	go.mongodb.org/mongo-driver v1.4.6 // indirect
	hotel-booking.com/v1/app v0.0.0-00010101000000-000000000000 // indirect
)
