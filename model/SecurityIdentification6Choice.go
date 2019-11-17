package model

// Choice between ISIN and an alternative format for the identification of a financial instrument. ISIN is the preferred format.
type SecurityIdentification6Choice struct {

	// International Securities Identification Number (ISIN).  A numbering system designed by the United Nation's International Organisation for Standardisation (ISO). The ISIN is composed of a 2-character prefix representing the country of issue, followed by the national security number (if one exists), and a check digit. Each country has a national numbering agency that assigns ISIN numbers for securities in that country.
	ISIN *ISINIdentifier `xml:"ISIN"`

	// Proprietary identification of a security assigned by an institution or organisation.
	OtherIdentification *AlternateSecurityIdentification1 `xml:"OthrId"`

	// Provides the ability to describe the instrument through a description and main characteristics.
	InstrumentDescription *SecurityInstrumentDescription2 `xml:"InstrmDesc"`
}

func (s *SecurityIdentification6Choice) SetISIN(value string) {
	s.ISIN = (*ISINIdentifier)(&value)
}

func (s *SecurityIdentification6Choice) AddOtherIdentification() *AlternateSecurityIdentification1 {
	s.OtherIdentification = new(AlternateSecurityIdentification1)
	return s.OtherIdentification
}

func (s *SecurityIdentification6Choice) AddInstrumentDescription() *SecurityInstrumentDescription2 {
	s.InstrumentDescription = new(SecurityInstrumentDescription2)
	return s.InstrumentDescription
}
