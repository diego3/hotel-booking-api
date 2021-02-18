package core

// Bed represents a bedroom
type Bed struct {
	size string
}

// Room represents a hotel room
type Room struct {
	code string
	beds []Bed
}
