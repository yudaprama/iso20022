package model

// Choice of format for the sub-balance.
type SubBalanceType5Choice struct {

	// Reason a security is not available or additional information about the financial instrument for which the balance is given, for example, unregistered, registered in nominee name.
	Code *SecuritiesBalanceType12Code `xml:"Cd"`

	// Reason a security is not available or additional information about the financial instrument for which the balance is given, for example, unregistered, registered in nominee name.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (s *SubBalanceType5Choice) SetCode(value string) {
	s.Code = (*SecuritiesBalanceType12Code)(&value)
}

func (s *SubBalanceType5Choice) AddProprietary() *GenericIdentification20 {
	s.Proprietary = new(GenericIdentification20)
	return s.Proprietary
}
