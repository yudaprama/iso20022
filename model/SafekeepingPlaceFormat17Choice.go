package model

// Choice between formats for the place of safekeeping.
type SafekeepingPlaceFormat17Choice struct {

	// Place of safekeeping expressed as a code and a narrative description.
	Identification *SafekeepingPlaceTypeAndText15 `xml:"Id"`

	// Place of safekeeping expressed with a country code.
	Country *CountryCode `xml:"Ctry"`

	// Place of safekeeping expressed with a type and identification.
	TypeAndIdentification *SafekeepingPlaceTypeAndAnyBICIdentifier1 `xml:"TpAndId"`

	// Place of safekeeping expressed with a propriety identification scheme.
	Proprietary *GenericIdentification85 `xml:"Prtry"`
}

func (s *SafekeepingPlaceFormat17Choice) AddIdentification() *SafekeepingPlaceTypeAndText15 {
	s.Identification = new(SafekeepingPlaceTypeAndText15)
	return s.Identification
}

func (s *SafekeepingPlaceFormat17Choice) SetCountry(value string) {
	s.Country = (*CountryCode)(&value)
}

func (s *SafekeepingPlaceFormat17Choice) AddTypeAndIdentification() *SafekeepingPlaceTypeAndAnyBICIdentifier1 {
	s.TypeAndIdentification = new(SafekeepingPlaceTypeAndAnyBICIdentifier1)
	return s.TypeAndIdentification
}

func (s *SafekeepingPlaceFormat17Choice) AddProprietary() *GenericIdentification85 {
	s.Proprietary = new(GenericIdentification85)
	return s.Proprietary
}
