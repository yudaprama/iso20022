package model

// Choice between formats for the place of safekeeping.
type SafekeepingPlaceFormat7Choice struct {

	// Place of safekeeping expressed as a code and a narrative description.
	Identification *SafekeepingPlaceTypeAndText1 `xml:"Id"`

	// Place of safekeeping expressed with a country code.
	Country *CountryCode `xml:"Ctry"`

	// Place of safekeeping expressed with a type and identification.
	TypeAndIdentification *SafekeepingPlaceTypeAndAnyBICIdentifier1 `xml:"TpAndId"`

	// Place of safekeeping expressed with a propriety identification scheme.
	Proprietary *GenericIdentification58 `xml:"Prtry"`
}

func (s *SafekeepingPlaceFormat7Choice) AddIdentification() *SafekeepingPlaceTypeAndText1 {
	s.Identification = new(SafekeepingPlaceTypeAndText1)
	return s.Identification
}

func (s *SafekeepingPlaceFormat7Choice) SetCountry(value string) {
	s.Country = (*CountryCode)(&value)
}

func (s *SafekeepingPlaceFormat7Choice) AddTypeAndIdentification() *SafekeepingPlaceTypeAndAnyBICIdentifier1 {
	s.TypeAndIdentification = new(SafekeepingPlaceTypeAndAnyBICIdentifier1)
	return s.TypeAndIdentification
}

func (s *SafekeepingPlaceFormat7Choice) AddProprietary() *GenericIdentification58 {
	s.Proprietary = new(GenericIdentification58)
	return s.Proprietary
}
