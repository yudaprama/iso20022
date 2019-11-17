package model

// Account to or from which a securities entry is made.
type SecuritiesAccount19 struct {

	// Unambiguous identification for the account between the account owner and the account servicer.
	Identification *Max35Text `xml:"Id"`

	// Specifies the type of securities account.
	Type *GenericIdentification30 `xml:"Tp,omitempty"`

	// Description of the account.
	Name *Max70Text `xml:"Nm,omitempty"`
}

func (s *SecuritiesAccount19) SetIdentification(value string) {
	s.Identification = (*Max35Text)(&value)
}

func (s *SecuritiesAccount19) AddType() *GenericIdentification30 {
	s.Type = new(GenericIdentification30)
	return s.Type
}

func (s *SecuritiesAccount19) SetName(value string) {
	s.Name = (*Max70Text)(&value)
}
