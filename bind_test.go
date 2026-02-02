package sp

import (
	"net/url"
	"slices"
	"testing"
)

const urlq = `searchableAttributes=title&attributesForFaceting=tags,authors,series,narrators&index=default&poot="toot"`

type params struct {
	SrchAttr  []string `qs:"searchableAttributes"`
	FacetAttr []string `qs:"attributesForFaceting"`
	Index     string   `qs:"index"`
	Poot      string   `qs:"-"`
}

var testParams = &params{
	SrchAttr:  []string{"title"},
	FacetAttr: []string{"tags", "authors", "series", "narrators"},
	Index:     "default",
	Poot:      "toot",
}

func TestUnmarshal(t *testing.T) {
	q := parsed()
	p := params{}
	err := Decode(q, &p)
	if err != nil {
		t.Error(err)
	}
	sw := []string{"title"}
	if !slices.Equal(p.SrchAttr, sw) {
		t.Errorf("got %v, expected %v\n", p.SrchAttr, sw)
	}
	facets := []string{"tags,authors,series,narrators"}
	if !slices.Equal(p.FacetAttr, facets) {
		t.Errorf("got %v, expected %v\n", p.FacetAttr, facets)
	}
	i := "default"
	if p.Index != i {
		t.Errorf("got %v, expected %v\n", p.Index, i)
	}
}

func TestMarshal(t *testing.T) {
	v, err := Encode(testParams)
	if err != nil {
		t.Error(err)
	}
	//fmt.Printf("%#v\n", v)
	sw := []string{"title"}
	if !slices.Equal(v["searchableAttributes"], sw) {
		t.Errorf("got %v, expected %v\n", v["searchableAttributes"], sw)
	}
	facets := []string{"tags", "authors", "series", "narrators"}
	if !slices.Equal(v["attributesForFaceting"], facets) {
		t.Errorf("got %v, expected %v\n", v["attributesForFaceting"], facets)
	}
	i := []string{"default"}
	if !slices.Equal(v["index"], i) {
		t.Errorf("got %v, expected %v\n", v["index"], i)
	}
}

func parsed() url.Values {
	v, _ := url.ParseQuery(urlq)
	//fmt.Printf("%#v\n", v)
	return v
}
