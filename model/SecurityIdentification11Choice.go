package model

// Choice between ISIN and an alternative format for the identification of a security. ISIN is the preferred format.
type SecurityIdentification11Choice struct {

	// International Securities Identification Number (ISIN).  A numbering system designed by the United Nation's International Organisation for Standardisation (ISO). The ISIN is composed of a 2-character prefix representing the country of issue, followed by the national security number (if one exists), and a check digit. Each country has a national numbering agency that assigns ISIN numbers for securities in that country.
	ISIN *ISINIdentifier `xml:"ISIN"`

	// Identification of a security by proprietary or domestic identification scheme.
	OtherIdentification *AlternateIdentification1 `xml:"OthrId"`
}

func (s *SecurityIdentification11Choice) SetISIN(value string) {
	s.ISIN = (*ISINIdentifier)(&value)
}

func (s *SecurityIdentification11Choice) AddOtherIdentification() *AlternateIdentification1 {
	s.OtherIdentification = new(AlternateIdentification1)
	return s.OtherIdentification
}
