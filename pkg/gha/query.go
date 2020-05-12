package gha

import "encoding/xml"

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
	XMLName             xml.Name            `xml:"Query"`
	Checkin             string              `xml:"Checkin,omitempty"`
	Nights              int                 `xml:"Nights,omitempty"`
	PropertyList        PropertyList        `xml:"PropertyList,omitempty"`
	HotelInfoProperties HotelInfoProperties `xml:"HotelInfoProperties,omitempty"`
}

// One or more IDs for hotel that require pricing updates.
type PropertyList struct {
	XMLName  xml.Name   `xml:"PropertyList"`
	Property []Property `xml:"Property,omitempty"`
}

// One or more properties for which Google wants updated room and Room Bundle
// metadata in a metadata Query message. This element can contain one or more
// <Property> elements that specify hotel property IDs.
type HotelInfoProperties struct {
	XMLName  xml.Name   `xml:"HotelInfoProperties"`
	Property []Property `xml:"Property,omitempty"`
}
