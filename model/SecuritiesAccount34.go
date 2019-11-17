package model

// Account to or from which a securities entry is made.
type SecuritiesAccount34 struct {

	// Unambiguous identification for the account between the account owner and the account servicer.”
	Identification *RestrictedFINXMax35Text `xml:"Id"`

	// Specifies the type of securities account.
	Type *PurposeCode8Choice `xml:"Tp,omitempty"`

	// Description of the account.
	Name *Max70Text `xml:"Nm,omitempty"`
}

func (s *SecuritiesAccount34) SetIdentification(value string) {
	s.Identification = (*RestrictedFINXMax35Text)(&value)
}

func (s *SecuritiesAccount34) AddType() *PurposeCode8Choice {
	s.Type = new(PurposeCode8Choice)
	return s.Type
}

func (s *SecuritiesAccount34) SetName(value string) {
	s.Name = (*Max70Text)(&value)
}
