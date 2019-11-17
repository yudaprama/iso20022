package model

// Choice of format for the sub-balance.
type SubBalanceType11Choice struct {

	// Reason a security is not available or additional information about the financial instrument for which the balance is given, for example, unregistered, registered in nominee name.
	Code *SecuritiesBalanceType12Code `xml:"Cd"`

	// Reason a security is not available or additional information about the financial instrument for which the balance is given, for example, unregistered, registered in nominee name.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (s *SubBalanceType11Choice) SetCode(value string) {
	s.Code = (*SecuritiesBalanceType12Code)(&value)
}

func (s *SubBalanceType11Choice) AddProprietary() *GenericIdentification30 {
	s.Proprietary = new(GenericIdentification30)
	return s.Proprietary
}
