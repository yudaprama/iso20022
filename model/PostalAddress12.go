package model

// Address of a party expressed in a formal structure.
type PostalAddress12 struct {

	// Name of a built-up area, with defined boundaries, and a local government.
	//
	TownName *Max35Text `xml:"TwnNm,omitempty"`

	// Identifies a subdivision of a country eg, state, region, county.
	CountrySubDivision *Max35Text `xml:"CtrySubDvsn,omitempty"`

	// Nation with its own government.
	Country *CountryCode `xml:"Ctry"`
}

func (p *PostalAddress12) SetTownName(value string) {
	p.TownName = (*Max35Text)(&value)
}

func (p *PostalAddress12) SetCountrySubDivision(value string) {
	p.CountrySubDivision = (*Max35Text)(&value)
}

func (p *PostalAddress12) SetCountry(value string) {
	p.Country = (*CountryCode)(&value)
}
