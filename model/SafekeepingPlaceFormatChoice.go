package model

// Choice of formats for the place of safekeeping.
type SafekeepingPlaceFormatChoice struct {

	// Place of safekeeping expressed as a code and BIC.
	Identification *SafekeepingPlaceAsCodeAndPartyIdentification `xml:"Id"`

	// Place of safekeeping expressed with a propriety identification scheme.
	IdentificationAsDSS *GenericIdentification5 `xml:"IdAsDSS"`

	// Place of safekeeping expressed with a country code.
	IdentificationAsCountry *CountryCode `xml:"IdAsCtry"`
}

func (s *SafekeepingPlaceFormatChoice) AddIdentification() *SafekeepingPlaceAsCodeAndPartyIdentification {
	s.Identification = new(SafekeepingPlaceAsCodeAndPartyIdentification)
	return s.Identification
}

func (s *SafekeepingPlaceFormatChoice) AddIdentificationAsDSS() *GenericIdentification5 {
	s.IdentificationAsDSS = new(GenericIdentification5)
	return s.IdentificationAsDSS
}

func (s *SafekeepingPlaceFormatChoice) SetIdentificationAsCountry(value string) {
	s.IdentificationAsCountry = (*CountryCode)(&value)
}
