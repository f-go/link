package gha

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"testing"
	"time"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestStructureOfHintRequest(t *testing.T) {
	example := `<?xml version="1.0" encoding="UTF-8"?>
<HintRequest id="request" timestamp="2019-06-03T22:59:48Z">
  <LastFetchTime>2019-06-03T22:54:40Z</LastFetchTime>
</HintRequest>`

	t.Run("Test HintRequest Struct", func(t *testing.T) {
		timestamp, _ := time.Parse(time.RFC3339, "2019-06-03T22:59:48Z")
		lastFetchTime, _ := time.Parse(time.RFC3339, "2019-06-03T22:54:40Z")
		want := HintRequest{
			XMLName:      xml.Name{"", "HintRequest"},
			ID:            "request",
			Timestamp:     timestamp,
			LastFetchTime: lastFetchTime,
		}

		got := HintRequest{}
		xml.Unmarshal([]byte(example), &got)

		if !reflect.DeepEqual(got.XMLName, want.XMLName) {
			t.Errorf("XMLName = %v, want %v", got.XMLName, want.XMLName)
		}
		if got.ID != want.ID {
			t.Errorf("ID = %v, want %v", got.ID, want.ID)
		}
		if got.Timestamp != want.Timestamp {
			t.Errorf("LastFetchTime = %v, want %v", got.Timestamp, want.Timestamp)
		}
		if got.LastFetchTime != want.LastFetchTime {
			t.Errorf("LastFetchTime = %v, want %v", got.LastFetchTime, want.LastFetchTime)
		}
	})
}

func printError(t *testing.T, name string, got, want interface{}) {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(fmt.Sprintf("%v", got), fmt.Sprintf("%v", want), false)
	t.Errorf("%s - Item\ngot:  %v\nwant: %v\ndiff: %v", name, got, want, dmp.DiffPrettyText(diffs))
}

func TestStructureOfHint(t *testing.T) {
	tests := []struct {
		name    string
		request string
		want    Hint
	}{
		{
			"Test Hint Struct - No Items",
			`<?xml version="1.0" encoding="UTF-8"?>
<Hint></Hint>`,
			Hint{
				xml.Name{"", "Hint"},
				[]Item{},
			},
		},

		// -------------------------------------------------------------------------------------------------------------
		{
			"Test Hint Struct - Exact Itinerary Hint Response",
			`<!-- Exact Itinerary Hint Response -->
<?xml version="1.0" encoding="UTF-8"?>
<Hint>
 <Item>
   <Property>12345</Property>
   <Stay>
     <CheckInDate>2018-07-03</CheckInDate>
     <LengthOfStay>3</LengthOfStay>
   </Stay>
 </Item>
 <Item>
   <Property>12345</Property>
   <Stay>
     <CheckInDate>2018-07-03</CheckInDate>
     <LengthOfStay>4</LengthOfStay>
   </Stay>
 </Item>
</Hint>`,
			Hint{
				xml.Name{"", "Hint"},
				[]Item{
					{
						XMLName:  xml.Name{"", "Item"},
						Property: []Property{{xml.Name{"", "Property"}, "12345"}},
						Stay: &Stay{
							XMLName:      xml.Name{"", "Stay"},
							CheckInDate:  "2018-07-03",
							LengthOfStay: 3,
						},
					},
					{
						XMLName:  xml.Name{"", "Item"},
						Property: []Property{{xml.Name{"", "Property"}, "12345"}},
						Stay: &Stay{
							XMLName:      xml.Name{"", "Stay"},
							CheckInDate:  "2018-07-03",
							LengthOfStay: 4,
						},
					},
				},
			},
		},

		// -------------------------------------------------------------------------------------------------------------
		{
			"Test Hint Struct - Check-in Ranges Hint Response",
			`<!-- Check-in Ranges Hint Response -->
<?xml version="1.0" encoding="UTF-8"?>
<Hint>
 <Item>
   <Property>12345</Property>
   <Property>67890</Property>
   <FirstDate>2018-07-03</FirstDate>
   <LastDate>2018-07-06</LastDate>
 </Item>
</Hint>`,
			Hint{
				xml.Name{"", "Hint"},
				[]Item{
					{
						XMLName: xml.Name{"", "Item"},
						Property: []Property{
							{xml.Name{"", "Property"}, "12345"},
							{xml.Name{"", "Property"}, "67890"},
						},
						FirstDate: "2018-07-03",
						LastDate:  "2018-07-06",
					},
				},
			},
		},

		// -------------------------------------------------------------------------------------------------------------
		{
			"Test Hint Struct - Ranged Stay Hint Response",
			`<!-- Ranged Stay Hint Response -->
<?xml version="1.0" encoding="UTF-8"?>
<Hint>
 <!-- Google fetches prices for all itineraries (first and last date are set) -->
 <Item>
   <Property>12345</Property>
   <StaysIncludingRange>
     <FirstDate>2018-07-03</FirstDate>
     <LastDate>2018-07-06</LastDate>
   </StaysIncludingRange>
 </Item>

 <!-- Google fetches prices for a single night (first date only) -->
 <Item>
   <Property>67890</Property>
   <StaysIncludingRange>
     <FirstDate>2018-07-03</FirstDate>
   </StaysIncludingRange>
 </Item>
</Hint>`,
			Hint{
				xml.Name{"", "Hint"},
				[]Item{
					{
						XMLName: xml.Name{"", "Item"},
						Property: []Property{
							{xml.Name{"", "Property"}, "12345"},
						},
						StaysIncludingRange: &StaysIncludingRange{
							XMLName:   xml.Name{"", "StaysIncludingRange"},
							FirstDate: "2018-07-03",
							LastDate:  "2018-07-06",
						},
					},
					{
						XMLName: xml.Name{"", "Item"},
						Property: []Property{
							{xml.Name{"", "Property"}, "67890"},
						},
						StaysIncludingRange: &StaysIncludingRange{
							XMLName:   xml.Name{"", "StaysIncludingRange"},
							FirstDate: "2018-07-03",
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Hint
			xml.Unmarshal([]byte(tt.request), &got)

			if !reflect.DeepEqual(got.XMLName, tt.want.XMLName) {
				printError(t, tt.name, got.XMLName, tt.want.XMLName)
			}
			if len(got.Item) != len(tt.want.Item) {
				printError(t, tt.name, len(got.Item), len(tt.want.Item))
				return // stop here if no of items do not match, to avoid index issues in the following loop
			}

			// check items
			for idx, item := range tt.want.Item {
				if !reflect.DeepEqual(item.Property, got.Item[idx].Property) {
					printError(t, tt.name, item.Property, got.Item[idx].Property)
				}
				if item.Stay != nil && !reflect.DeepEqual(item.Stay, got.Item[idx].Stay) {
					printError(t, tt.name, item.Stay, got.Item[idx].Stay)
				}
				if item.StaysIncludingRange != nil && !reflect.DeepEqual(item.StaysIncludingRange, got.Item[idx].StaysIncludingRange) {
					printError(t, tt.name, item.StaysIncludingRange, got.Item[idx].StaysIncludingRange)
				}
				if item.FirstDate != got.Item[idx].FirstDate {
					printError(t, tt.name, item.FirstDate, got.Item[idx].FirstDate)
				}
				if item.LastDate != got.Item[idx].LastDate {
					printError(t, tt.name, item.LastDate, got.Item[idx].LastDate)
				}
			}
		})
	}
}
