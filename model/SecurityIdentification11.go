package model

// Identification of a security.
type SecurityIdentification11 struct {

	// Identification of a security.
	Identification *SecurityIdentification11Choice `xml:"Id"`

	// Textual description of a security instrument.
	Description *Max140Text `xml:"Desc,omitempty"`
}

func (s *SecurityIdentification11) AddIdentification() *SecurityIdentification11Choice {
	s.Identification = new(SecurityIdentification11Choice)
	return s.Identification
}

func (s *SecurityIdentification11) SetDescription(value string) {
	s.Description = (*Max140Text)(&value)
}
