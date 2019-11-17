package model

// Choice of formats for the place of safekeeping.
type SafekeepingPlaceFormat26Choice struct {

	// Place of safekeeping expressed as a code and a narrative description.
	Identification *SafekeepingPlaceTypeAndText25 `xml:"Id"`

	// Place of safekeeping expressed with a country code.
	Country *CountryCode `xml:"Ctry"`

	// Place of safekeeping expressed with a type and identification.
	TypeAndIdentification *SafekeepingPlaceTypeAndAnyBICIdentifier1 `xml:"TpAndId"`

	// Place of safekeeping expressed with a propriety identification scheme.
	Proprietary *GenericIdentification89 `xml:"Prtry"`
}

func (s *SafekeepingPlaceFormat26Choice) AddIdentification() *SafekeepingPlaceTypeAndText25 {
	s.Identification = new(SafekeepingPlaceTypeAndText25)
	return s.Identification
}

func (s *SafekeepingPlaceFormat26Choice) SetCountry(value string) {
	s.Country = (*CountryCode)(&value)
}

func (s *SafekeepingPlaceFormat26Choice) AddTypeAndIdentification() *SafekeepingPlaceTypeAndAnyBICIdentifier1 {
	s.TypeAndIdentification = new(SafekeepingPlaceTypeAndAnyBICIdentifier1)
	return s.TypeAndIdentification
}

func (s *SafekeepingPlaceFormat26Choice) AddProprietary() *GenericIdentification89 {
	s.Proprietary = new(GenericIdentification89)
	return s.Proprietary
}
