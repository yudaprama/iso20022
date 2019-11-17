package model

// Account to or from which a securities entry is made.
type SecuritiesAccount13 struct {

	// Unambiguous identification for the account between the account owner and the account servicer.”
	Identification *Max35Text `xml:"Id"`

	// Specifies the type of securities account.
	Type *GenericIdentification20 `xml:"Tp,omitempty"`

	// Description of the account.
	Name *Max70Text `xml:"Nm,omitempty"`
}

func (s *SecuritiesAccount13) SetIdentification(value string) {
	s.Identification = (*Max35Text)(&value)
}

func (s *SecuritiesAccount13) AddType() *GenericIdentification20 {
	s.Type = new(GenericIdentification20)
	return s.Type
}

func (s *SecuritiesAccount13) SetName(value string) {
	s.Name = (*Max70Text)(&value)
}
