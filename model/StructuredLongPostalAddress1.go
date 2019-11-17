package model

// Address of a party expressed in a formal structure, usually according to the country's postal services specifications.
type StructuredLongPostalAddress1 struct {

	// Name of the building or house.
	BuildingName *Max35Text `xml:"BldgNm,omitempty"`

	// Name of a street or thoroughfare.
	StreetName *Max35Text `xml:"StrtNm,omitempty"`

	// Number that identifies the position of a building on a street.
	StreetBuildingIdentification *Max35Text `xml:"StrtBldgId,omitempty"`

	// Floor or storey within a building.
	Floor *Max16Text `xml:"Flr,omitempty"`

	// Name of a built-up area, with defined boundaries, and a local government.
	TownName *Max35Text `xml:"TwnNm"`

	// Name of a district, ie, a part of a town or region.
	DistrictName *Max35Text `xml:"DstrctNm,omitempty"`

	// Identification of an administrative division of a country, state, or territory.
	RegionIdentification *Max35Text `xml:"RgnId,omitempty"`

	// Organised political community or area forming a part of a federation.
	State *Max35Text `xml:"Stat,omitempty"`

	// Identifier of a county.
	CountyIdentification *Max35Text `xml:"CtyId,omitempty"`

	// Nation with its own government.
	Country *CountryCode `xml:"Ctry"`

	// Identifier consisting of a group of letters and/or numbers that is added to a postal address to assist the sorting of mail.
	PostCodeIdentification *Max16Text `xml:"PstCdId"`

	// Numbered box in a post office, assigned to a person or organisation, where letters are kept until called for.
	PostOfficeBox *Max16Text `xml:"POB,omitempty"`
}

func (s *StructuredLongPostalAddress1) SetBuildingName(value string) {
	s.BuildingName = (*Max35Text)(&value)
}

func (s *StructuredLongPostalAddress1) SetStreetName(value string) {
	s.StreetName = (*Max35Text)(&value)
}

func (s *StructuredLongPostalAddress1) SetStreetBuildingIdentification(value string) {
	s.StreetBuildingIdentification = (*Max35Text)(&value)
}

func (s *StructuredLongPostalAddress1) SetFloor(value string) {
	s.Floor = (*Max16Text)(&value)
}

func (s *StructuredLongPostalAddress1) SetTownName(value string) {
	s.TownName = (*Max35Text)(&value)
}

func (s *StructuredLongPostalAddress1) SetDistrictName(value string) {
	s.DistrictName = (*Max35Text)(&value)
}

func (s *StructuredLongPostalAddress1) SetRegionIdentification(value string) {
	s.RegionIdentification = (*Max35Text)(&value)
}

func (s *StructuredLongPostalAddress1) SetState(value string) {
	s.State = (*Max35Text)(&value)
}

func (s *StructuredLongPostalAddress1) SetCountyIdentification(value string) {
	s.CountyIdentification = (*Max35Text)(&value)
}

func (s *StructuredLongPostalAddress1) SetCountry(value string) {
	s.Country = (*CountryCode)(&value)
}

func (s *StructuredLongPostalAddress1) SetPostCodeIdentification(value string) {
	s.PostCodeIdentification = (*Max16Text)(&value)
}

func (s *StructuredLongPostalAddress1) SetPostOfficeBox(value string) {
	s.PostOfficeBox = (*Max16Text)(&value)
}
