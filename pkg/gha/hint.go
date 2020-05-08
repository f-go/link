package gha

import (
	"encoding/xml"
	"time"
)

// Defines a Hint Request message that contains the time Google last
// received an update from your server.
type HintRequest struct {
	XMLName       xml.Name  `xml:"HintRequest"`
	ID            string    `xml:"id,attr"`
	Timestamp     time.Time `xml:"timestamp,attr"`
	LastFetchTime time.Time `xml:"LastFetchTime"`
}

// Hint Response message that specifies the hotels whose prices have
// changed since the last time Google received a successful Hint Response
// from those same servers.
type Hint struct {
	XMLName xml.Name `xml:"Hint"`
	Item    []Item   `xml:"Item"`
}

// A container for the hotel/itinerary to be updated.
type Item struct {
	XMLName             xml.Name             `xml:"Item"`
	Property            []Property           `xml:"Property"`
	FirstDate           string               `xml:"FirstDate,omitempty"`
	LastDate            string               `xml:"LastDate,omitempty"`
	Stay                *Stay                `xml:"Stay,omitempty"`
	StaysIncludingRange *StaysIncludingRange `xml:"StaysIncludingRange,omitempty"`
}

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
	XMLName xml.Name `xml:"Property"`
	ID      string   `xml:",chardata"`
}

// A container for the checkin date and length of stay elements in  an exact
// itinerary Hint Response message. Each <Item> can contain only a single <Stay>.
type Stay struct {
	XMLName      xml.Name `xml:"Stay"`
	CheckInDate  string   `xml:"CheckInDate"`
	LengthOfStay int8     `xml:"LengthOfStay"`
}

// A container for the first date and last date elements in a ranged stay
// Hint Response message.
type StaysIncludingRange struct {
	XMLName   xml.Name `xml:"StaysIncludingRange"`
	FirstDate string   `xml:"FirstDate"`
	LastDate  string   `xml:"LastDate"`
}
