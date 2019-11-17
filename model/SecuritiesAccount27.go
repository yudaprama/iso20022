package model

// Account to or from which a securities entry is made.
type SecuritiesAccount27 struct {

	// Unambiguous identification for the account between the account owner and the account servicer.”
	Identification *RestrictedFINXMax35Text `xml:"Id"`

	// Specifies the type of securities account.
	Type *GenericIdentification47 `xml:"Tp,omitempty"`

	// Description of the account.
	Name *Max70Text `xml:"Nm,omitempty"`
}

func (s *SecuritiesAccount27) SetIdentification(value string) {
	s.Identification = (*RestrictedFINXMax35Text)(&value)
}

func (s *SecuritiesAccount27) AddType() *GenericIdentification47 {
	s.Type = new(GenericIdentification47)
	return s.Type
}

func (s *SecuritiesAccount27) SetName(value string) {
	s.Name = (*Max70Text)(&value)
}
