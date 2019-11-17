package model

// Account to or from which a securities entry is made.
type SecuritiesAccount25 struct {

	// Unambiguous identification for the account between the account owner and the account servicer.”
	Identification *Max35Text `xml:"Id"`

	// Specifies the type of securities account.
	Type *PurposeCode7Choice `xml:"Tp,omitempty"`

	// Description of the account.
	Name *Max70Text `xml:"Nm,omitempty"`
}

func (s *SecuritiesAccount25) SetIdentification(value string) {
	s.Identification = (*Max35Text)(&value)
}

func (s *SecuritiesAccount25) AddType() *PurposeCode7Choice {
	s.Type = new(PurposeCode7Choice)
	return s.Type
}

func (s *SecuritiesAccount25) SetName(value string) {
	s.Name = (*Max70Text)(&value)
}
