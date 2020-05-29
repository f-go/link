package gha

import (
	"encoding/xml"

	cdt "github.com/f-go/go-custom-datetime"
)

// Query messages are requests from Google for pricing or
// metadata updates. They are used with both the Pull and
// Pull with Hints delivery modes.
//
// There are two types of Query messages:
//
// - Pricing:
//   Google requests pricing updates for the specified hotels.
//   When you receive a pricing Query message, you should respond
//   with a <Transaction> message that contains the requested
//   pricing information in <Result> elements.
//
//   Live Queries are a special type of pricing Query message in
//   which Google asks for real-time price updates.
//
//   For more information, consult Pricing Overview.
//
// - Metadata:
//   Google requests metadata updates for the rooms and Room Bundles
//   for the specified hotels. When you receive a metadata Query message,
//   you should respond with a <Transaction> message that specifies data
//   about the rooms and Room Bundles in <PropertyDataSet> elements.
//
//   For more information, consult Room Bundle metadata.
//
// The syntax for the messages is different, depending on the type. Both
// types are described in this section.
type Query struct {
	XMLName             xml.Name             `xml:""`
	Checkin             cdt.CustomDate      `xml:",omitempty"`
	Nights              int                  `xml:",omitempty"`
	PropertyList        *PropertyList        `xml:",omitempty"`
	HotelInfoProperties *HotelInfoProperties `xml:",omitempty"`

	// Left out Live Query support (<LatencySensitive> and <Context>) for now.
	// This will be implemented later.
	//
	// https://developers.google.com/hotels/hotel-prices/xml-reference/queries#Context
}

// One or more IDs for hotel that require pricing updates.
type PropertyList struct {
	XMLName  xml.Name   `xml:""`
	Property []Property `xml:",omitempty"`
}

// One or more properties for which Google wants updated room and Room Bundle
// metadata in a metadata Query message. This element can contain one or more
// <Property> elements that specify hotel property IDs.
type HotelInfoProperties struct {
	XMLName  xml.Name   `xml:""`
	Property []Property `xml:",omitempty"`
}
