package gha

import (
	"encoding/xml"

	cdt "github.com/f-go/go-custom-datetime"
)

// Defines a Hint Request message that contains the time Google last
// received an update from your server.
type HintRequest struct {
	XMLName       xml.Name           `xml:""`
	ID            string             `xml:"id,attr"`
	Timestamp     cdt.CustomDateTime `xml:"timestamp,attr"`
	LastFetchTime cdt.CustomDateTime `xml:""`
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
	FirstDate           cdt.CustomDate       `xml:",omitempty"`
	LastDate            cdt.CustomDate       `xml:",omitempty"`
	Stay                *Stay                `xml:",omitempty"`
	StaysIncludingRange *StaysIncludingRange `xml:",omitempty"`
}

// A container for the checkin date and length of stay elements in  an exact
// itinerary Hint Response message. Each <Item> can contain only a single <Stay>.
type Stay struct {
	XMLName      xml.Name       `xml:""`
	CheckInDate  cdt.CustomDate `xml:""`
	LengthOfStay int8           `xml:""`
}

// A container for the first date and last date elements in a ranged stay
// Hint Response message.
type StaysIncludingRange struct {
	XMLName   xml.Name       `xml:""`
	FirstDate cdt.CustomDate `xml:""`
	LastDate  cdt.CustomDate `xml:""`
}
