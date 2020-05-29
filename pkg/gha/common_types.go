package gha

// The ID of a hotel, using the same ID as the Hotel List Feed. The number of
// <Property> elements you can specify in a single <Item> block is determined
// by the type of Hint Response message:
// - Exact itineraries:
//   Up to 100 hotels.
// - Check-in ranges:
//   More than one if you set <MultipleItineraries> to "checkin_range" in your
//   <QueryControl> message.
// - Ranged stay:
//   More than one if you set <MultipleItineraries> to "affected_dates"
//   in your <QueryControl> message.
type Property struct {
	ID string `xml:",chardata"`
}

// Represents an amount of money with its currency type.
type Money struct {
	Value    float32 `xml:",chardata"`
	Currency string  `xml:"currency,attr"`
}
