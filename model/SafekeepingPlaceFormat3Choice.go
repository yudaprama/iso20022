package model

// Choice between formats for the place of safekeeping.
type SafekeepingPlaceFormat3Choice struct {

	// Place of safekeeping expressed as a code and a narrative description.
	Identification *SafekeepingPlaceTypeAndText3 `xml:"Id"`

	// Place of safekeeping expressed with a country code.
	Country *CountryCode `xml:"Ctry"`

	// Place of safekeeping expressed with a type and identification.
	TypeAndIdentification *SafekeepingPlaceTypeAndAnyBICIdentifier1 `xml:"TpAndId"`

	// Place of safekeeping expressed with a propriety identification scheme.
	Proprietary *GenericIdentification21 `xml:"Prtry"`
}

func (s *SafekeepingPlaceFormat3Choice) AddIdentification() *SafekeepingPlaceTypeAndText3 {
	s.Identification = new(SafekeepingPlaceTypeAndText3)
	return s.Identification
}

func (s *SafekeepingPlaceFormat3Choice) SetCountry(value string) {
	s.Country = (*CountryCode)(&value)
}

func (s *SafekeepingPlaceFormat3Choice) AddTypeAndIdentification() *SafekeepingPlaceTypeAndAnyBICIdentifier1 {
	s.TypeAndIdentification = new(SafekeepingPlaceTypeAndAnyBICIdentifier1)
	return s.TypeAndIdentification
}

func (s *SafekeepingPlaceFormat3Choice) AddProprietary() *GenericIdentification21 {
	s.Proprietary = new(GenericIdentification21)
	return s.Proprietary
}
