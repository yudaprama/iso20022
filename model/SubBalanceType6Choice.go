package model

// Choice of format for the sub-balance.
type SubBalanceType6Choice struct {

	// Reason a security is not available or additional information about the financial instrument for which the balance is given, for example, unregistered, registered in nominee name.
	Code *SecuritiesBalanceType7Code `xml:"Cd"`

	// Reason a security is not available or additional information about the financial instrument for which the balance is given, for example, unregistered, registered in nominee name.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (s *SubBalanceType6Choice) SetCode(value string) {
	s.Code = (*SecuritiesBalanceType7Code)(&value)
}

func (s *SubBalanceType6Choice) AddProprietary() *GenericIdentification20 {
	s.Proprietary = new(GenericIdentification20)
	return s.Proprietary
}
