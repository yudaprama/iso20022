package model

// Choice between ISIN and an alternative format for the identification of a security. ISIN is the preferred format.
type SecurityIdentification7 struct {

	// International Securities Identification Number (ISIN).  A numbering system designed by the United Nation's International Organisation for Standardisation (ISO). The ISIN is composed of a 2-character prefix representing the country of issue, followed by the national security number (if one exists), and a check digit. Each country has a national numbering agency that assigns ISIN numbers for securities in that country.
	ISIN *ISINIdentifier `xml:"ISIN"`

	// Proprietary identification of a security assigned by an institution or organisation.
	OtherIdentification *AlternateSecurityIdentification3 `xml:"OthrId"`

	// Textual description of a security instrument.
	Description *Max140Text `xml:"Desc,omitempty"`
}

func (s *SecurityIdentification7) SetISIN(value string) {
	s.ISIN = (*ISINIdentifier)(&value)
}

func (s *SecurityIdentification7) AddOtherIdentification() *AlternateSecurityIdentification3 {
	s.OtherIdentification = new(AlternateSecurityIdentification3)
	return s.OtherIdentification
}

func (s *SecurityIdentification7) SetDescription(value string) {
	s.Description = (*Max140Text)(&value)
}
