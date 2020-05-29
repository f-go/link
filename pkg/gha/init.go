package gha

import (
	"time"

	cdt "github.com/f-go/go-custom-datetime"
)

// Initializes the default date and time formats required by the
// Google Hotels API.
//
// https://developers.google.com/hotels/hotel-prices/xml-reference/datetime
func init() {
	cdt.CustomDateFormat = "2006-01-02"
	cdt.CustomTimeFormat = "15:04"
	cdt.CustomDateTimeFormat = time.RFC3339
}
