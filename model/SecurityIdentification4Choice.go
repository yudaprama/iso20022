package model

// Choice between formats for the identification of a financial instrument.
type SecurityIdentification4Choice struct {

	// International Securities Identification Number (ISIN).  A numbering system designed by the United Nation's International Organisation for Standardisation (ISO). The ISIN is composed of a 2-character prefix representing the country of issue, followed by the national security number (if one exists), and a check digit. Each country has a national numbering agency that assigns ISIN numbers for securities in that country.
	ISIN *ISINIdentifier `xml:"ISIN"`

	// Proprietary identification of an underlying financial instrument.
	Proprietary *AlternateSecurityIdentification2 `xml:"Prtry"`
}

func (s *SecurityIdentification4Choice) SetISIN(value string) {
	s.ISIN = (*ISINIdentifier)(&value)
}

func (s *SecurityIdentification4Choice) AddProprietary() *AlternateSecurityIdentification2 {
	s.Proprietary = new(AlternateSecurityIdentification2)
	return s.Proprietary
}
