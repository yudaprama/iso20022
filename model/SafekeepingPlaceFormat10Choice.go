package model

// Choice between formats for the place of safekeeping.
type SafekeepingPlaceFormat10Choice struct {

	// Place of safekeeping expressed as a code and a narrative description.
	Identification *SafekeepingPlaceTypeAndText8 `xml:"Id"`

	// Place of safekeeping expressed with a country code.
	Country *CountryCode `xml:"Ctry"`

	// Place of safekeeping expressed with a type and identification.
	TypeAndIdentification *SafekeepingPlaceTypeAndAnyBICIdentifier1 `xml:"TpAndId"`

	// Place of safekeeping expressed with a propriety identification scheme.
	Proprietary *GenericIdentification78 `xml:"Prtry"`
}

func (s *SafekeepingPlaceFormat10Choice) AddIdentification() *SafekeepingPlaceTypeAndText8 {
	s.Identification = new(SafekeepingPlaceTypeAndText8)
	return s.Identification
}

func (s *SafekeepingPlaceFormat10Choice) SetCountry(value string) {
	s.Country = (*CountryCode)(&value)
}

func (s *SafekeepingPlaceFormat10Choice) AddTypeAndIdentification() *SafekeepingPlaceTypeAndAnyBICIdentifier1 {
	s.TypeAndIdentification = new(SafekeepingPlaceTypeAndAnyBICIdentifier1)
	return s.TypeAndIdentification
}

func (s *SafekeepingPlaceFormat10Choice) AddProprietary() *GenericIdentification78 {
	s.Proprietary = new(GenericIdentification78)
	return s.Proprietary
}
