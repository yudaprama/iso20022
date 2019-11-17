package model

// Address of a party expressed in a formal structure, usually according to the country's postal services specifications.
type PostalAddress5 struct {

	// Name of a street or thoroughfare.
	StreetName *Max70Text `xml:"StrtNm,omitempty"`

	// Identifier consisting of a group of letters and/or numbers that is added to a postal address to assist the sorting of mail.
	PostCodeIdentification *Max16Text `xml:"PstCdId,omitempty"`

	// Name of a built-up area, with defined boundaries, and a local government.
	TownName *Max35Text `xml:"TwnNm,omitempty"`

	// Identifies a subdivision of a country eg, state, region, county.
	CountrySubDivision *Max35Text `xml:"CtrySubDvsn,omitempty"`

	// Nation with its own government.
	Country *CountryCode `xml:"Ctry"`
}

func (p *PostalAddress5) SetStreetName(value string) {
	p.StreetName = (*Max70Text)(&value)
}

func (p *PostalAddress5) SetPostCodeIdentification(value string) {
	p.PostCodeIdentification = (*Max16Text)(&value)
}

func (p *PostalAddress5) SetTownName(value string) {
	p.TownName = (*Max35Text)(&value)
}

func (p *PostalAddress5) SetCountrySubDivision(value string) {
	p.CountrySubDivision = (*Max35Text)(&value)
}

func (p *PostalAddress5) SetCountry(value string) {
	p.Country = (*CountryCode)(&value)
}
