package gha

import (
	"encoding/xml"
	"testing"
)

func TestPropertyStruct(t *testing.T) {
	example := "<Property>abc</Property>"

	var p Property
	if err := xml.Unmarshal([]byte(example), &p); err != nil {
		t.Errorf("Unmarshal data failed. %v", err)
	}

	if p.ID != "abc" {
		t.Errorf("ID got %v, want: abc", p.ID)
	}
}

func TestMoneyStruct(t *testing.T) {
	example := "<Money currency=\"USD\">13.54</Money>"

	var m Money
	if err := xml.Unmarshal([]byte(example), &m); err != nil {
		t.Errorf("Unmarshal data failed. %v", err)
	}

	if m.Value != 13.54 {
		t.Errorf("Value got %v, want: 13.54", m.Value)
	}
	if m.Currency != "USD" {
		t.Errorf("ID got %v, want: USD", m.Currency)
	}
}
