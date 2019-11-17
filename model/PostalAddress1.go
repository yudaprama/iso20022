package model

// Information that locates and identifies a specific address, as defined by postal services.
type PostalAddress1 struct {

	// Identifies the nature of the postal address.
	AddressType *AddressType2Code `xml:"AdrTp,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services, that is presented in free format text.
	AddressLine []*Max70Text `xml:"AdrLine,omitempty"`

	// Name of a street or thoroughfare.
	StreetName *Max70Text `xml:"StrtNm,omitempty"`

	// Number that identifies the position of a building on a street.
	BuildingNumber *Max16Text `xml:"BldgNb,omitempty"`

	// Identifier consisting of a group of letters and/or numbers that is added to a postal address to assist the sorting of mail.
	PostCode *Max16Text `xml:"PstCd,omitempty"`

	// Name of a built-up area, with defined boundaries, and a local government.
	TownName *Max35Text `xml:"TwnNm,omitempty"`

	// Identifies a subdivision of a country for example, state, region, county.
	CountrySubDivision *Max35Text `xml:"CtrySubDvsn,omitempty"`

	// Nation with its own government.
	Country *CountryCode `xml:"Ctry"`
}

func (p *PostalAddress1) SetAddressType(value string) {
	p.AddressType = (*AddressType2Code)(&value)
}

func (p *PostalAddress1) AddAddressLine(value string) {
	p.AddressLine = append(p.AddressLine, (*Max70Text)(&value))
}

func (p *PostalAddress1) SetStreetName(value string) {
	p.StreetName = (*Max70Text)(&value)
}

func (p *PostalAddress1) SetBuildingNumber(value string) {
	p.BuildingNumber = (*Max16Text)(&value)
}

func (p *PostalAddress1) SetPostCode(value string) {
	p.PostCode = (*Max16Text)(&value)
}

func (p *PostalAddress1) SetTownName(value string) {
	p.TownName = (*Max35Text)(&value)
}

func (p *PostalAddress1) SetCountrySubDivision(value string) {
	p.CountrySubDivision = (*Max35Text)(&value)
}

func (p *PostalAddress1) SetCountry(value string) {
	p.Country = (*CountryCode)(&value)
}
