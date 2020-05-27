package gha

import (
	"encoding/xml"
	"time"
)

// Container for descriptive information about rooms and packages
// and/or pricing and availability for rooms and packages.
//
// Transaction messages can have any number of child elements,
// as long as the total message size does not exceed 100 MB.
//
// * At least one of <PropertyDataSet> or <Result> is required.
// * Note:
//   HTML syntax is not allowed within your XML elements (even if
//   it's escaped). All Transaction messages containing HTML will
//   be rejected.
type Transaction struct {
	XMLName   xml.Name  `xml:"Transaction"`
	ID        string    `xml:"id,attr"`
	Timestamp time.Time `xml:"timestamp,attr"`
	Partner   string    `xml:"partner,attr,omitempty"`
	Result    []Result  `xml:"Result,omitempty"`

	// PropertyDataSet PropertyDataSet `xml:"PropertyDataSet,omitempty"` will be added later

}

// Container for pricing and availability updates in a <Transaction> message.
//
// Pricing data for a room's itinerary or a <RoomBundle> element that
// defines Room Bundles and additional types of rooms for the property.
// The <Result> element can also be used to remove itineraries from
// inventory.
//
// * Note:
//   Leave <BaseRate> for the <Result> empty for shared rooms (e.g., dorm-style
//   hostels). Instead, convey shared/hostel pricing info in <RoomBundle>.
//
// * The Rates field:
//   Used only when there are multiple rates for the same room/itinerary
//   combination. For example, you define multiple rates for conditional
//   rates, qualified rates, or conditional rates in Room Bundles.
//
// https://developers.google.com/hotels/hotel-prices/xml-reference/transaction-messages#Result
type Result struct {
	Rate

	XMLName  xml.Name `xml:""`
	Property Property `xml:""`
	Checkin  string   `xml:""`
	RoomID   string   `xml:",omitempty"`
	Nights   uint8    `xml:",omitempty"`
	Rates    *Rates   `xml:",omitempty"`
}

// Container for one or more <Rate> blocks. Each <Rate> in <Rates>
// defines a different price for the room/itinerary combination.
type Rates struct {
	XMLName xml.Name `xml:""`
	Rate    []Rate   `xml:",omitempty"`
}

// Container that holds additional information such as the number
// and type of guests (adults or children).
type OccupancyDetails struct {
	XMLName   xml.Name  `xml:""`
	NumAdults uint8     `xml:""` // min:1, max:20
	Children  *Children `xml:",omitempty"`
}

// Container that holds a list of the maximum age for each child.
type Children struct {
	XMLName xml.Name `xml:""`
	Child   []Child  `xml:",omitempty"`
}

// Specifies which guests are children (typically age 0-17).
type Child struct {
	XMLName xml.Name `xml:""`
	Age     uint8    `xml:"age,attr"`
}

// Container fot one or more landing pages that are eligible for the hotel.
//
// A landing page is a website that can handle the booking process for the end-user.
// To explicitly include certain landing page (and exclude others), add one or more
// <AllowablePointsOfSale> elements that match the <PointOfSale> element's id
// attribute in the landing pages file.
//
// If you do not include this element, all landing pages defined in the landing pages
// file are considered eligible to be used for booking the room. For more information,
// refer to Landing Pages File Syntax.
//
// see also: https://developers.google.com/hotels/hotel-prices/dev-guide/pos-syntax
type AllowablePointsOfSale struct {
	XMLName     xml.Name      `xml:""`
	PointOfSale []PointOfSale `xml:",omitempty"`
}

// Container for a point of sale.
//
// see also: https://developers.google.com/hotels/hotel-prices/dev-guide/pos-syntax
type PointOfSale struct {
	XMLName xml.Name `xml:""`
	ID      string   `xml:"id,attr"`
}

// Container that holds refund information.
//
// Enables listing a rate as being fully refundable or providing a free cancellation.
// If not provided, no information about a refund is displayed. A refund policy at the
// <PackageData> level overrides the refund policy at the <Result> level. A refund
// policy at the <Rates> level overrides the refund policy at the <PackageData> level.
// Refundable pricing can also be highlighted to users through alternative options
// without directly modifying your transaction message schema. Learn more about these
// options here: https://support.google.com/hotelprices/answer/9824483
//
// Note:
// We recommend setting all of the attributes. A feed status warning message is
// generated when one or more attributes are not set.
//
// If you do not set any attributes, the rate does not display as refundable.
// The attributes are:
// * available: (Required)
//   Set to 1 or true to indicate if the rate allows a full refund; otherwise set
//   to 0 or false.
// * refundable_until_days: (Required if available is true)
//   Specifies the number of days in advance of check-in that a full refund can be
//   requested. The value of refundable_until_days must be an integer between 0 and
//   330, inclusive.
// * refundable_until_time: (Highly recommended if available is true) Specifies the
//   latest time of day, in the local time of the hotel, that a full refund request
//   will be honored. This can be combined with refundable_until_days to specify,
//   for example, that "refunds are available until 4:00PM two days before check-in".
//   If refundable_until_time isn't set, the value defaults to midnight.
//
//   The value of this attribute uses the Time format.
//
// When setting the attributes, note the following:
// * If available or refundable_until_days isn't set, the rate does not display as
//   refundable.
// * If available is 0 or false, the other attributes are ignored. The rate does not
//   display as refundable even if one or both of the other attributes is set.
type Refundable struct {
	XMLName             xml.Name `xml:""`
	Available           bool     `xml:"available,attr"`
	RefundableUntilDays int32    `xml:"refundable_until_days,attr"`
	RefundableUntilTime string   `xml:"refundable_until_time,attr"`
}

// Container that holds all rate information.
//
// Values set in a <Rate> override pricing-related values on the parent <Result>
// or <RoomBundle> element. If they are not set in <Rate>, they inherit their
// value from the parent element.
type Rate struct {
	XMLName               xml.Name               `xml:""`
	RateRuleID            string                 `xml:"rate_rule_id,attr,omitempty"`
	Baserate              *Money                 `xml:",omitempty"`
	Tax                   *Money                 `xml:",omitempty"`
	OtherFees             *Money                 `xml:",omitempty"`
	ExpirationTime        *time.Time             `xml:",omitempty"`
	Refundable            *Refundable            `xml:",omitempty"`
	ChargeCurrency        string                 `xml:",omitempty"` // [deposit|hotel|installment|web]
	AllowablePointsOfSale *AllowablePointsOfSale `xml:",omitempty"`
	Occupancy             uint8                  `xml:",omitempty"`
	OccupancyDetails      *OccupancyDetails      `xml:",omitempty"`
	Custom1               string                 `xml:",omitempty"`
	Custom2               string                 `xml:",omitempty"`
	Custom3               string                 `xml:",omitempty"`
	Custom4               string                 `xml:",omitempty"`
	Custom5               string                 `xml:",omitempty"`
}
