package gha

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"
	"time"

	cdt "github.com/f-go/go-custom-datetime"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestHintRequestStruct(t *testing.T) {
	request, err := ioutil.ReadFile("./testdata/HintRequest.xml")
	if err != nil {
		t.Errorf("File reading error %v", err)
		return
	}

	t.Run("Test HintRequest Struct", func(t *testing.T) {
		timestamp, _ := cdt.NewCustomDateTime("2019-06-03T22:59:48Z")
		lastFetchTime, _ := cdt.NewCustomDateTime("2019-06-03T22:54:40Z")
		want := HintRequest{
			XMLName:       xml.Name{Local: "HintRequest"},
			ID:            "request",
			Timestamp:     timestamp,
			LastFetchTime: lastFetchTime,
		}

		got := HintRequest{}
		err := xml.Unmarshal(request, &got)
		if err != nil {
			t.Errorf("Parsing request data failed. %v", err)
		}

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

func printError(t *testing.T, got, want interface{}) {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(fmt.Sprintf("%v", got), fmt.Sprintf("%v", want), false)
	t.Errorf("\ngot:  %v\nwant: %v\ndiff: %v", got, want, dmp.DiffPrettyText(diffs))
}

func newCustomDate(value string) cdt.CustomDate {
	d, _ := cdt.NewCustomDate(value)
	return d
}

func TestHintStruct(t *testing.T) {
	tests := []struct {
		name string
		file string
		want Hint
	}{
		{
			"Test Hint Struct - No Items",
			"./testdata/Hint-NoItems.xml",
			Hint{
				xml.Name{Local: "Hint"},
				[]Item{},
			},
		},

		// -------------------------------------------------------------------------------------------------------------
		{
			"Test Hint Struct - Exact Itinerary Hint Response",
			"./testdata/Hint-ExactItinerary.xml",
			Hint{
				xml.Name{Local: "Hint"},
				[]Item{
					{
						XMLName:  xml.Name{Local: "Item"},
						Property: []Property{{xml.Name{Local: "Property"}, "12345"}},
						Stay: &Stay{
							XMLName:      xml.Name{Local: "Stay"},
							CheckInDate:  newCustomDate("2018-07-03"),
							LengthOfStay: 3,
						},
					},
					{
						XMLName:  xml.Name{Local: "Item"},
						Property: []Property{{xml.Name{Local: "Property"}, "12345"}},
						Stay: &Stay{
							XMLName:      xml.Name{Local: "Stay"},
							CheckInDate:  newCustomDate("2018-07-03"),
							LengthOfStay: 4,
						},
					},
				},
			},
		},

		// -------------------------------------------------------------------------------------------------------------
		{
			"Test Hint Struct - Check-in Ranges Hint Response",
			"./testdata/Hint-CheckInRanges.xml",
			Hint{
				xml.Name{Local: "Hint"},
				[]Item{
					{
						XMLName: xml.Name{Local: "Item"},
						Property: []Property{
							{xml.Name{Local: "Property"}, "12345"},
							{xml.Name{Local: "Property"}, "67890"},
						},
						FirstDate: newCustomDate("2018-07-03"),
						LastDate:  newCustomDate("2018-07-06"),
					},
				},
			},
		},

		// -------------------------------------------------------------------------------------------------------------
		{
			"Test Hint Struct - Ranged Stay Hint Response",
			"./testdata/Hint-RangedStay.xml",
			Hint{
				xml.Name{Local: "Hint"},
				[]Item{
					{
						XMLName: xml.Name{Local: "Item"},
						Property: []Property{
							{xml.Name{Local: "Property"}, "12345"},
						},
						StaysIncludingRange: &StaysIncludingRange{
							XMLName:   xml.Name{Local: "StaysIncludingRange"},
							FirstDate: newCustomDate("2018-07-03"),
							LastDate:  newCustomDate("2018-07-06"),
						},
					},
					{
						XMLName: xml.Name{Local: "Item"},
						Property: []Property{
							{xml.Name{Local: "Property"}, "67890"},
						},
						StaysIncludingRange: &StaysIncludingRange{
							XMLName:   xml.Name{Local: "StaysIncludingRange"},
							FirstDate: newCustomDate("2018-07-03"),
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request, err := ioutil.ReadFile(tt.file)
			if err != nil {
				t.Errorf("File reading error %v", err)
				return
			}

			var got Hint
			err = xml.Unmarshal(request, &got)
			if err != nil {
				t.Errorf("%s: Parsing request data failed with error: %v", tt.name, err)
				return
			}

			if !reflect.DeepEqual(got.XMLName, tt.want.XMLName) {
				printError(t, got.XMLName, tt.want.XMLName)
			}
			if len(got.Item) != len(tt.want.Item) {
				printError(t, len(got.Item), len(tt.want.Item))
				return // stop here if no of items do not match, to avoid index issues in the following loop
			}

			// check items
			for idx, item := range tt.want.Item {
				if !reflect.DeepEqual(item.Property, got.Item[idx].Property) {
					printError(t, item.Property, got.Item[idx].Property)
				}
				if item.Stay != nil && !reflect.DeepEqual(item.Stay, got.Item[idx].Stay) {
					printError(t, item.Stay, got.Item[idx].Stay)
				}
				if item.StaysIncludingRange != nil && !reflect.DeepEqual(item.StaysIncludingRange, got.Item[idx].StaysIncludingRange) {
					printError(t, item.StaysIncludingRange, got.Item[idx].StaysIncludingRange)
				}
				if !reflect.DeepEqual(item.FirstDate, got.Item[idx].FirstDate) {
					printError(
						t,
						time.Time(item.FirstDate).Format(cdt.CustomDateFormat),
						time.Time(got.Item[idx].FirstDate).Format(cdt.CustomDateFormat))
				}
				if !reflect.DeepEqual(item.LastDate, got.Item[idx].LastDate) {
					printError(
						t,
						time.Time(item.LastDate).Format(cdt.CustomDateFormat),
						time.Time(got.Item[idx].LastDate).Format(cdt.CustomDateFormat))
				}
			}
		})
	}
}
