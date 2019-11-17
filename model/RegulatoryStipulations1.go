package model

// Specifies regulatory stipulations that financial institutions must be compliant with in the country, region, and/or area they conduct business.
type RegulatoryStipulations1 struct {

	// Nation with its own government, occupying a particular territory.
	Country *CountryCode `xml:"Ctry"`

	// Specifies regulatory stipulations that financial institutions must be compliant with in the country, region, and/or where they conduct business.
	Stipulations []*Max350Text `xml:"Stiptns"`
}

func (r *RegulatoryStipulations1) SetCountry(value string) {
	r.Country = (*CountryCode)(&value)
}

func (r *RegulatoryStipulations1) AddStipulations(value string) {
	r.Stipulations = append(r.Stipulations, (*Max350Text)(&value))
}
