package gha

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"

	cdt "github.com/f-go/go-custom-datetime"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestTransactionMultiPropertyExample(t *testing.T) {
	var got, want Transaction

	request, err := ioutil.ReadFile("./testdata/Transaction-MultiPropertyExample.xml")
	if err != nil {
		t.Errorf("File reading error %v", err)
		return
	}

	if err = xml.Unmarshal(request, &got); err != nil {
		t.Errorf("Parsing request data failed with error: %v", err)
		return
	}

	timestamp, _ := cdt.NewCustomDateTime("2017-07-23T16:20:00-04:00")
	checkin, _ := cdt.NewCustomDate("2018-06-10")
	want = Transaction{
		XMLName:   xml.Name{Local: "Transaction"},
		ID:        "42",
		Timestamp: cdt.CustomDateTime(timestamp),
		Result: []Result{
			{
				XMLName:  xml.Name{Local: "Result"},
				Property: Property{xml.Name{Local: "Property"}, "060773"},
				RoomID:   "RoomType101",
				Checkin:  checkin,
				Nights:   2,
				Rate: Rate{
					Baserate:  &Money{278.33, "USD"},
					Tax:       &Money{25.12, "USD"},
					OtherFees: &Money{2.00, "USD"},
					AllowablePointsOfSale: &AllowablePointsOfSale{
						XMLName: xml.Name{Local: "AllowablePointsOfSale"},
						PointOfSale: []PointOfSale{
							{xml.Name{Local: "PointOfSale"}, "site1"},
						},
					},
				},
			},
			{
				XMLName:  xml.Name{Local: "Result"},
				Property: Property{xml.Name{Local: "Property"}, "052213"},
				RoomID:   "RoomType101",
				Checkin:  checkin,
				Nights:   2,
				Rate: Rate{
					Baserate:  &Money{299.98, "USD"},
					Tax:       &Money{26.42, "USD"},
					OtherFees: &Money{2.00, "USD"},
					AllowablePointsOfSale: &AllowablePointsOfSale{
						XMLName: xml.Name{Local: "AllowablePointsOfSale"},
						PointOfSale: []PointOfSale{
							{xml.Name{Local: "PointOfSale"}, "otto"},
							{xml.Name{Local: "PointOfSale"}, "simon"},
						},
					},
				},
			},
		},
	}

	if !reflect.DeepEqual(got.Result, want.Result) {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(fmt.Sprintf("%v", got), fmt.Sprintf("%v", want), false)
		t.Errorf("Query\ngot:  %v\nwant: %v\ndiff: %v", got, want, dmp.DiffPrettyText(diffs))
	}
}

func TestTransactionMultiRateExample(t *testing.T) {
	var got, want Transaction

	request, err := ioutil.ReadFile("./testdata/Transaction-MultiRateExample.xml")
	if err != nil {
		t.Errorf("File reading error %v", err)
		return
	}

	if err = xml.Unmarshal(request, &got); err != nil {
		t.Errorf("Parsing request data failed with error: %v", err)
		return
	}

	timestamp, _ := cdt.NewCustomDateTime("2020-07-23T16:20:00-04:00")
	checkin, _ := cdt.NewCustomDate("2021-01-13")
	want = Transaction{
		XMLName:   xml.Name{Local: "Transaction"},
		ID:        "42",
		Timestamp: cdt.CustomDateTime(timestamp),
		Result: []Result{
			{
				XMLName:  xml.Name{Local: "Result"},
				Property: Property{xml.Name{Local: "Property"}, "1234"},
				Checkin:  checkin,
				Nights:   9,
				Rate: Rate{
					Baserate:  &Money{3196.1, "USD"},
					Tax:       &Money{559.49, "USD"},
					OtherFees: &Money{543.34, "USD"},
					Occupancy: 2,
				},
				Rates: &Rates{
					XMLName: xml.Name{Local: "Rates"},
					Rate: []Rate{
						{
							XMLName:   xml.Name{Local: "Rate"},
							Baserate:  &Money{3196.1, "USD"},
							Tax:       &Money{559.49, "USD"},
							OtherFees: &Money{543.34, "USD"},
							Occupancy: 1,
						},
						{
							XMLName:   xml.Name{Local: "Rate"},
							Baserate:  &Money{3196.1, "USD"},
							Tax:       &Money{559.49, "USD"},
							OtherFees: &Money{543.34, "USD"},
							Occupancy: 3,
						},
						{
							XMLName:   xml.Name{Local: "Rate"},
							Baserate:  &Money{3196.1, "USD"},
							Tax:       &Money{559.49, "USD"},
							OtherFees: &Money{543.34, "USD"},
							Occupancy: 4,
						},
						{
							XMLName:   xml.Name{Local: "Rate"},
							Baserate:  &Money{3196.1, "USD"},
							Tax:       &Money{559.49, "USD"},
							OtherFees: &Money{543.34, "USD"},
							Occupancy: 5,
						},
						{
							XMLName:   xml.Name{Local: "Rate"},
							Baserate:  &Money{3196.1, "USD"},
							Tax:       &Money{559.49, "USD"},
							OtherFees: &Money{543.34, "USD"},
							Occupancy: 6,
						},
					},
				},
			},
		},
	}

	if !reflect.DeepEqual(got, want) {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(fmt.Sprintf("%v", got), fmt.Sprintf("%v", want), false)
		t.Errorf("Query\ngot:  %v\nwant: %v\ndiff: %v", got, want, dmp.DiffPrettyText(diffs))
	}
}

func TestTransactionBaseRateAndConditionalRate(t *testing.T) {
	var got, want Transaction

	request, err := ioutil.ReadFile("./testdata/Transaction-BaseRateAndConditionalRate.xml")
	if err != nil {
		t.Errorf("File reading error %v", err)
		return
	}

	if err = xml.Unmarshal(request, &got); err != nil {
		t.Errorf("Parsing request data failed with error: %v", err)
		return
	}

	timestamp, _ := cdt.NewCustomDateTime("2017-07-18T16:20:00-04:00")
	checkin, _ := cdt.NewCustomDate("2018-06-10")
	want = Transaction{
		XMLName:   xml.Name{Local: "Transaction"},
		ID:        "42",
		Timestamp: cdt.CustomDateTime(timestamp),
		Result: []Result{
			{
				XMLName:  xml.Name{Local: "Result"},
				Property: Property{xml.Name{Local: "Property"}, "1234"},
				Checkin:  checkin,
				Nights:   1,
				Rate: Rate{
					Baserate:  &Money{200.00, "USD"},
					Tax:       &Money{20.00, "USD"},
					OtherFees: &Money{1.00, "USD"},
				},
				Rates: &Rates{
					XMLName: xml.Name{Local: "Rates"},
					Rate: []Rate{
						{
							XMLName:    xml.Name{Local: "Rate"},
							RateRuleID: "mobile",
							Baserate:   &Money{180.00, "USD"},
							Tax:        &Money{18.00, "USD"},
							Custom1:    "ratecode123",
						},
					},
				},
			},
		},
	}

	if !reflect.DeepEqual(got, want) {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(fmt.Sprintf("%v", got), fmt.Sprintf("%v", want), false)
		t.Errorf("Query\ngot:  %v\nwant: %v\ndiff: %v", got, want, dmp.DiffPrettyText(diffs))
	}
}

func TestTransactionOneItineraryPricingForOneAdultChild(t *testing.T) {
	var got, want Transaction

	request, err := ioutil.ReadFile("./testdata/Transaction-OneItineraryPricingForOneAdultChild.xml")
	if err != nil {
		t.Errorf("File reading error %v", err)
		return
	}

	if err = xml.Unmarshal(request, &got); err != nil {
		t.Errorf("Parsing request data failed with error: %v", err)
		return
	}

	timestamp, _ := cdt.NewCustomDateTime("2018-04-18T11:27:45-04:00")
	checkin, _ := cdt.NewCustomDate("2018-06-20")
	want = Transaction{
		XMLName:   xml.Name{Local: "Transaction"},
		ID:        "Wtdj8QoQIWcAAbaTGlIAAAC4",
		Timestamp: cdt.CustomDateTime(timestamp),
		Result: []Result{
			{
				XMLName:  xml.Name{Local: "Result"},
				Property: Property{xml.Name{Local: "Property"}, "8251"},
				Checkin:  checkin,
				Nights:   1,
				Rate: Rate{
					Baserate:  &Money{62.18, "USD"},
					Tax:       &Money{2.45, "USD"},
					OtherFees: &Money{0.00, "USD"},
				},
				Rates: &Rates{
					XMLName: xml.Name{Local: "Rates"},
					Rate: []Rate{
						{
							XMLName:    xml.Name{Local: "Rate"},
							RateRuleID: "rule-951",
							Occupancy:  2,
							OccupancyDetails: &OccupancyDetails{
								XMLName:   xml.Name{Local: "OccupancyDetails"},
								NumAdults: 1,
								Children: &Children{
									XMLName: xml.Name{Local: "Children"},
									Child: []Child{
										{xml.Name{Local: "Child"}, 17},
									},
								},
							},
							Baserate:  &Money{42.61, "USD"},
							Tax:       &Money{5.70, "USD"},
							OtherFees: &Money{0.00, "USD"},
							Custom1:   "abc4",
							AllowablePointsOfSale: &AllowablePointsOfSale{
								XMLName: xml.Name{Local: "AllowablePointsOfSale"},
								PointOfSale: []PointOfSale{
									{xml.Name{Local: "PointOfSale"}, "yourhotelpartnersite.com"},
								},
							},
						},
					},
				},
			},
		},
	}

	if !reflect.DeepEqual(got, want) {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(fmt.Sprintf("%v", got.Result), fmt.Sprintf("%v", want.Result), false)
		t.Errorf("Query\ngot:  %v\nwant: %v\ndiff: %v", got.Result, want.Result, dmp.DiffPrettyText(diffs))
	}
}
