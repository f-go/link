package gha

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"

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

	want = Query{
		XMLName: xml.Name{Local: "Query"},
		Checkin: "2018-06-10",
		Nights:  3,
		PropertyList: PropertyList{
			XMLName: xml.Name{Local: "PropertyList"},
			Property: []Property{
				{xml.Name{Local: "Property"}, "pid5"},
				{xml.Name{Local: "Property"}, "pid8"},
				{xml.Name{Local: "Property"}, "pid13"},
				{xml.Name{Local: "Property"}, "pid21"},
			},
		},
	}

	data, _ := xml.Marshal(want)
	fmt.Println(string(data))

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
		XMLName: xml.Name{Local: "Query"},
		HotelInfoProperties: HotelInfoProperties{
			XMLName: xml.Name{Local: "HotelInfoProperties"},
			Property: []Property{
				{xml.Name{Local: "Property"}, "pid5"},
				{xml.Name{Local: "Property"}, "pid8"},
				{xml.Name{Local: "Property"}, "pid13"},
				{xml.Name{Local: "Property"}, "pid21"},
			},
		},
	}

	if !reflect.DeepEqual(got, want) {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(fmt.Sprintf("%v", got), fmt.Sprintf("%v", want), false)
		t.Errorf("Query\ngot:  %v\nwant: %v\ndiff: %v", got, want, dmp.DiffPrettyText(diffs))
	}
}
