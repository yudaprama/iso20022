package model

// Choice of formats for the place of safekeeping.
type SafekeepingPlaceFormat27Choice struct {

	// Place of safekeeping expressed as a code and a narrative description.
	Identification *SafekeepingPlaceTypeAndText9 `xml:"Id"`

	// Place of safekeeping expressed with a country code.
	Country *CountryCode `xml:"Ctry"`

	// Place of safekeeping expressed with a type and identification.
	TypeAndIdentification *SafekeepingPlaceTypeAndAnyBICIdentifier1 `xml:"TpAndId"`

	// Place of safekeeping expressed with a propriety identification scheme.
	Proprietary *GenericIdentification89 `xml:"Prtry"`
}

func (s *SafekeepingPlaceFormat27Choice) AddIdentification() *SafekeepingPlaceTypeAndText9 {
	s.Identification = new(SafekeepingPlaceTypeAndText9)
	return s.Identification
}

func (s *SafekeepingPlaceFormat27Choice) SetCountry(value string) {
	s.Country = (*CountryCode)(&value)
}

func (s *SafekeepingPlaceFormat27Choice) AddTypeAndIdentification() *SafekeepingPlaceTypeAndAnyBICIdentifier1 {
	s.TypeAndIdentification = new(SafekeepingPlaceTypeAndAnyBICIdentifier1)
	return s.TypeAndIdentification
}

func (s *SafekeepingPlaceFormat27Choice) AddProprietary() *GenericIdentification89 {
	s.Proprietary = new(GenericIdentification89)
	return s.Proprietary
}
