package model

// Account to or from which a securities entry is made.
type SecuritiesAccount18 struct {

	// Unambiguous identification for the account between the account owner and the account servicer.
	Identification *Max35Text `xml:"Id"`

	// Specifies if the account is a House, a Client or a Liquidity Provider (Market Maker) account.
	Type *ClearingAccountType1Code `xml:"Tp"`

	// Description of the account.
	Name *Max70Text `xml:"Nm,omitempty"`
}

func (s *SecuritiesAccount18) SetIdentification(value string) {
	s.Identification = (*Max35Text)(&value)
}

func (s *SecuritiesAccount18) SetType(value string) {
	s.Type = (*ClearingAccountType1Code)(&value)
}

func (s *SecuritiesAccount18) SetName(value string) {
	s.Name = (*Max70Text)(&value)
}
