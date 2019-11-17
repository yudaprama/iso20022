package model

// Number assigned by a government agency to identify foreign nationals.
type SecuritiesAccount3 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *Max35Text `xml:"Id"`

	// Choice between a code and a data source scheme to identify the type of account.
	Type *PurposeCode5Choice `xml:"Tp,omitempty"`

	// .
	Name *Max70Text `xml:"Nm,omitempty"`
}

func (s *SecuritiesAccount3) SetIdentification(value string) {
	s.Identification = (*Max35Text)(&value)
}

func (s *SecuritiesAccount3) AddType() *PurposeCode5Choice {
	s.Type = new(PurposeCode5Choice)
	return s.Type
}

func (s *SecuritiesAccount3) SetName(value string) {
	s.Name = (*Max70Text)(&value)
}
