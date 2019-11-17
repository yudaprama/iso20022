package model

// Choice of location of a residence of a party.
type ResidenceLocation1Choice struct {

	// Specifies the account owner's resident country.
	Country *CountryCode `xml:"Ctry"`

	// Specifies the account owner's resident geographical region or area.
	Area *Max35Text `xml:"Area"`
}

func (r *ResidenceLocation1Choice) SetCountry(value string) {
	r.Country = (*CountryCode)(&value)
}

func (r *ResidenceLocation1Choice) SetArea(value string) {
	r.Area = (*Max35Text)(&value)
}
