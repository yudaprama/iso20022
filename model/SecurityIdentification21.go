package model

// Identification of a security.
type SecurityIdentification21 struct {

	// International Securities Identification Number (ISIN).  A numbering system designed by the United Nation's International Organisation for Standardisation (ISO). The ISIN is composed of a 2-character prefix representing the country of issue, followed by the national security number (if one exists), and a check digit. Each country has a national numbering agency that assigns ISIN numbers for securities in that country.
	ISIN *ISINOct2015Identifier `xml:"ISIN,omitempty"`

	// Identification of a security by proprietary or domestic identification scheme.
	OtherIdentification []*OtherIdentification2 `xml:"OthrId,omitempty"`

	// Textual description of a security instrument.
	Description *RestrictedFINXMax140Text `xml:"Desc,omitempty"`
}

func (s *SecurityIdentification21) SetISIN(value string) {
	s.ISIN = (*ISINOct2015Identifier)(&value)
}

func (s *SecurityIdentification21) AddOtherIdentification() *OtherIdentification2 {
	newValue := new(OtherIdentification2)
	s.OtherIdentification = append(s.OtherIdentification, newValue)
	return newValue
}

func (s *SecurityIdentification21) SetDescription(value string) {
	s.Description = (*RestrictedFINXMax140Text)(&value)
}
