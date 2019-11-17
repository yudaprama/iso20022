package model

// Information that locates and identifies a specific address, as defined by postal services.
type PostalAddress21 struct {

	// Type of address.
	AddressType *AddressType2Choice `xml:"AdrTp,omitempty"`

	// Indicates whether mail should be sent to the address.
	MailingIndicator *YesNoIndicator `xml:"MlngInd,omitempty"`

	// Indicates whether the address is the official address of the party.
	RegistrationAddressIndicator *YesNoIndicator `xml:"RegnAdrInd,omitempty"`

	// When the individual resides at another person’s address, the name of the other person.
	CareOf *Max70Text `xml:"CareOf,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services, that is presented in free format text.
	AddressLine []*Max70Text `xml:"AdrLine,omitempty"`

	// Name of the street or thoroughfare.
	StreetName *Max70Text `xml:"StrtNm,omitempty"`

	// Number that identifies the position of the building on a street.
	BuildingNumber *Max16Text `xml:"BldgNb,omitempty"`

	// Name of the building or house.
	BuildingName *Max35Text `xml:"BldgNm,omitempty"`

	// Post box number of the addressee within the residential or company building.
	PostBox *Max10Text `xml:"PstBx,omitempty"`

	// Side or wing of the building, for example, ‘wing A’.
	SideInBuilding *Max35Text `xml:"SdInBldg,omitempty"`

	// Floor or storey within the building.
	Floor *Max70Text `xml:"Flr,omitempty"`

	// Identification of the suite or apartment.
	SuiteIdentification *Max10Text `xml:"SuiteId,omitempty"`

	// Identifier consisting of a group of letters and/or numbers that is added to a postal address to assist the sorting of mail.
	PostCode *Max16Text `xml:"PstCd,omitempty"`

	// Name of a district, that is, a part of the town or region.
	DistrictName *Max35Text `xml:"DstrctNm,omitempty"`

	// Name of the village.
	Village *Max70Text `xml:"Vllg,omitempty"`

	// Name of the town or city.
	TownName *Max35Text `xml:"TwnNm,omitempty"`

	// Name of the state, county or country sub-division.
	State *Max70Text `xml:"Stat,omitempty"`

	// Country of the address.
	Country *CountryCode `xml:"Ctry"`
}

func (p *PostalAddress21) AddAddressType() *AddressType2Choice {
	p.AddressType = new(AddressType2Choice)
	return p.AddressType
}

func (p *PostalAddress21) SetMailingIndicator(value string) {
	p.MailingIndicator = (*YesNoIndicator)(&value)
}

func (p *PostalAddress21) SetRegistrationAddressIndicator(value string) {
	p.RegistrationAddressIndicator = (*YesNoIndicator)(&value)
}

func (p *PostalAddress21) SetCareOf(value string) {
	p.CareOf = (*Max70Text)(&value)
}

func (p *PostalAddress21) AddAddressLine(value string) {
	p.AddressLine = append(p.AddressLine, (*Max70Text)(&value))
}

func (p *PostalAddress21) SetStreetName(value string) {
	p.StreetName = (*Max70Text)(&value)
}

func (p *PostalAddress21) SetBuildingNumber(value string) {
	p.BuildingNumber = (*Max16Text)(&value)
}

func (p *PostalAddress21) SetBuildingName(value string) {
	p.BuildingName = (*Max35Text)(&value)
}

func (p *PostalAddress21) SetPostBox(value string) {
	p.PostBox = (*Max10Text)(&value)
}

func (p *PostalAddress21) SetSideInBuilding(value string) {
	p.SideInBuilding = (*Max35Text)(&value)
}

func (p *PostalAddress21) SetFloor(value string) {
	p.Floor = (*Max70Text)(&value)
}

func (p *PostalAddress21) SetSuiteIdentification(value string) {
	p.SuiteIdentification = (*Max10Text)(&value)
}

func (p *PostalAddress21) SetPostCode(value string) {
	p.PostCode = (*Max16Text)(&value)
}

func (p *PostalAddress21) SetDistrictName(value string) {
	p.DistrictName = (*Max35Text)(&value)
}

func (p *PostalAddress21) SetVillage(value string) {
	p.Village = (*Max70Text)(&value)
}

func (p *PostalAddress21) SetTownName(value string) {
	p.TownName = (*Max35Text)(&value)
}

func (p *PostalAddress21) SetState(value string) {
	p.State = (*Max70Text)(&value)
}

func (p *PostalAddress21) SetCountry(value string) {
	p.Country = (*CountryCode)(&value)
}
