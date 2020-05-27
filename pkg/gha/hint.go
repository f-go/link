package gha

import (
	"encoding/xml"
	"time"
)

// Defines a Hint Request message that contains the time Google last
// received an update from your server.
type HintRequest struct {
	XMLName       xml.Name  `xml:""`
	ID            string    `xml:"id,attr"`
	Timestamp     time.Time `xml:"timestamp,attr"`
	LastFetchTime time.Time `xml:""`
}

// Hint Response message that specifies the hotels whose prices have
// changed since the last time Google received a successful Hint Response
// from those same servers.
type Hint struct {
	XMLName xml.Name `xml:""`
	Item    []Item   `xml:""`
}

// A container for the hotel/itinerary to be updated.
type Item struct {
	XMLName             xml.Name             `xml:""`
	Property            []Property           `xml:""`
	FirstDate           string               `xml:",omitempty"`
	LastDate            string               `xml:",omitempty"`
	Stay                *Stay                `xml:",omitempty"`
	StaysIncludingRange *StaysIncludingRange `xml:",omitempty"`
}

// A container for the checkin date and length of stay elements in  an exact
// itinerary Hint Response message. Each <Item> can contain only a single <Stay>.
type Stay struct {
	XMLName      xml.Name `xml:""`
	CheckInDate  string   `xml:""`
	LengthOfStay int8     `xml:""`
}

// A container for the first date and last date elements in a ranged stay
// Hint Response message.
type StaysIncludingRange struct {
	XMLName   xml.Name `xml:""`
	FirstDate string   `xml:""`
	LastDate  string   `xml:""`
}
