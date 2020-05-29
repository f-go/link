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

func TestPricingQuery(t *testing.T) {
	var got, want Query

	request, err := ioutil.ReadFile("./testdata/Query-PricingQuery.xml")
	if err != nil {
		t.Errorf("File reading error %v", err)
		return
	}

	if err = xml.Unmarshal(request, &got); err != nil {
		t.Errorf("Parsing request data failed with error: %v", err)
		return
	}

	checkin, _ := cdt.NewCustomDate("2018-06-10")
	want = Query{
		Checkin: checkin,
		Nights:  3,
		PropertyList: &PropertyList{
			Property: []Property{
				{"pid5"},
				{"pid8"},
				{"pid13"},
				{"pid21"},
			},
		},
	}

	if !reflect.DeepEqual(got, want) {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(fmt.Sprintf("%v", got), fmt.Sprintf("%v", want), false)
		t.Errorf("Query\ngot:  %v\nwant: %v\ndiff: %v", got, want, dmp.DiffPrettyText(diffs))
	}
}

func TestMetadataQuery(t *testing.T) {
	var got, want Query

	request, err := ioutil.ReadFile("./testdata/Query-MetadataQuery.xml")
	if err != nil {
		t.Errorf("File reading error %v", err)
		return
	}
	if err = xml.Unmarshal(request, &got); err != nil {
		t.Errorf("Parsing request data failed with error: %v", err)
		return
	}

	want = Query{
		HotelInfoProperties: &HotelInfoProperties{
			Property: []Property{
				{"pid5"},
				{"pid8"},
				{"pid13"},
				{"pid21"},
			},
		},
	}

	if !reflect.DeepEqual(got, want) {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(fmt.Sprintf("%v", got), fmt.Sprintf("%v", want), false)
		t.Errorf("Query\ngot:  %v\nwant: %v\ndiff: %v", got, want, dmp.DiffPrettyText(diffs))
	}
}
