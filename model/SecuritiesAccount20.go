package model

// Number assigned by a government agency to identify foreign nationals.
type SecuritiesAccount20 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *Max35Text `xml:"Id"`

	// Choice between a code and a data source scheme to identify the type of account.
	Type *ClearingAccountType1Code `xml:"Tp"`

	// .
	Name *Max70Text `xml:"Nm,omitempty"`
}

func (s *SecuritiesAccount20) SetIdentification(value string) {
	s.Identification = (*Max35Text)(&value)
}

func (s *SecuritiesAccount20) SetType(value string) {
	s.Type = (*ClearingAccountType1Code)(&value)
}

func (s *SecuritiesAccount20) SetName(value string) {
	s.Name = (*Max70Text)(&value)
}
