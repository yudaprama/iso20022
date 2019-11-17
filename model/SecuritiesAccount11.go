package model

// Account to or from which a securities entry is made.
type SecuritiesAccount11 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *Max35Text `xml:"Id"`

	// Specifies the type of securities account.
	Type *PurposeCode1Choice `xml:"Tp,omitempty"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max70Text `xml:"Nm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`
}

func (s *SecuritiesAccount11) SetIdentification(value string) {
	s.Identification = (*Max35Text)(&value)
}

func (s *SecuritiesAccount11) AddType() *PurposeCode1Choice {
	s.Type = new(PurposeCode1Choice)
	return s.Type
}

func (s *SecuritiesAccount11) SetName(value string) {
	s.Name = (*Max70Text)(&value)
}

func (s *SecuritiesAccount11) SetDesignation(value string) {
	s.Designation = (*Max35Text)(&value)
}
