package model

// Address of a party expressed in a formal structure, usually according to the country's postal services specifications.
type PostalAddress2 struct {

	// Name of a street or thoroughfare.
	StreetName *Max70Text `xml:"StrtNm,omitempty"`

	// Identifier consisting of a group of letters and/or numbers that is added to a postal address to assist the sorting of mail.
	PostCodeIdentification *Max16Text `xml:"PstCdId"`

	// Name of a built-up area, with defined boundaries, and a local government.
	TownName *Max35Text `xml:"TwnNm"`

	// Identifies a subdivision of a country for example, state, region, county.
	CountrySubDivision *Max35Text `xml:"CtrySubDvsn,omitempty"`

	// Nation with its own government.
	Country *CountryCode `xml:"Ctry"`
}

func (p *PostalAddress2) SetStreetName(value string) {
	p.StreetName = (*Max70Text)(&value)
}

func (p *PostalAddress2) SetPostCodeIdentification(value string) {
	p.PostCodeIdentification = (*Max16Text)(&value)
}

func (p *PostalAddress2) SetTownName(value string) {
	p.TownName = (*Max35Text)(&value)
}

func (p *PostalAddress2) SetCountrySubDivision(value string) {
	p.CountrySubDivision = (*Max35Text)(&value)
}

func (p *PostalAddress2) SetCountry(value string) {
	p.Country = (*CountryCode)(&value)
}
