package model

// Information that locates and identifies a specific address, as defined by postal services.
type PostalAddress11 struct {

	// Identifies the nature of the postal address.
	AddressType *AddressType2Code `xml:"AdrTp,omitempty"`

	// Identification of a division of a large organisation or building.
	Department *Max70Text `xml:"Dept,omitempty"`

	// Identification of a sub-division of a large organisation or building.
	SubDepartment *Max70Text `xml:"SubDept,omitempty"`

	// Name of a street or thoroughfare.
	StreetName *Max70Text `xml:"StrtNm,omitempty"`

	// Number that identifies the position of a building on a street.
	BuildingNumber *Max16Text `xml:"BldgNb,omitempty"`

	// Identifier consisting of a group of letters and/or numbers that is added to a postal address to assist the sorting of mail.
	PostCode *Max16Text `xml:"PstCd,omitempty"`

	// Name of a built-up area, with defined boundaries, and a local government.
	TownName *Max35Text `xml:"TwnNm,omitempty"`

	// Identifies a subdivision of a country such as state, region, county.
	CountrySubDivision *Max35Text `xml:"CtrySubDvsn,omitempty"`

	// Nation with its own government.
	Country *CountryCode `xml:"Ctry,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services, presented in free format text.
	AddressLine []*Max70Text `xml:"AdrLine,omitempty"`

	// Floor or storey within a building.
	Floor *Max70Text `xml:"Flr,omitempty"`

	// Numbered box in a post office, assigned to a person or organisation, where letters are kept until called for.
	PostBox *Max16Text `xml:"PstBx,omitempty"`

	// Name of the building or house.
	BuildingName *Max70Text `xml:"BldgNm,omitempty"`

	// Building room number.
	Room *Max70Text `xml:"Room,omitempty"`
}

func (p *PostalAddress11) SetAddressType(value string) {
	p.AddressType = (*AddressType2Code)(&value)
}

func (p *PostalAddress11) SetDepartment(value string) {
	p.Department = (*Max70Text)(&value)
}

func (p *PostalAddress11) SetSubDepartment(value string) {
	p.SubDepartment = (*Max70Text)(&value)
}

func (p *PostalAddress11) SetStreetName(value string) {
	p.StreetName = (*Max70Text)(&value)
}

func (p *PostalAddress11) SetBuildingNumber(value string) {
	p.BuildingNumber = (*Max16Text)(&value)
}

func (p *PostalAddress11) SetPostCode(value string) {
	p.PostCode = (*Max16Text)(&value)
}

func (p *PostalAddress11) SetTownName(value string) {
	p.TownName = (*Max35Text)(&value)
}

func (p *PostalAddress11) SetCountrySubDivision(value string) {
	p.CountrySubDivision = (*Max35Text)(&value)
}

func (p *PostalAddress11) SetCountry(value string) {
	p.Country = (*CountryCode)(&value)
}

func (p *PostalAddress11) AddAddressLine(value string) {
	p.AddressLine = append(p.AddressLine, (*Max70Text)(&value))
}

func (p *PostalAddress11) SetFloor(value string) {
	p.Floor = (*Max70Text)(&value)
}

func (p *PostalAddress11) SetPostBox(value string) {
	p.PostBox = (*Max16Text)(&value)
}

func (p *PostalAddress11) SetBuildingName(value string) {
	p.BuildingName = (*Max70Text)(&value)
}

func (p *PostalAddress11) SetRoom(value string) {
	p.Room = (*Max70Text)(&value)
}
