package model

// Choice of format for the sub-balance.
type SubBalanceType13Choice struct {

	// Reason a security is not available or additional information about the financial instrument for which the balance is given, for example, unregistered, registered in nominee name.
	Code *SecuritiesBalanceType12Code `xml:"Cd"`

	// Reason a security is not available or additional information about the financial instrument for which the balance is given, for example, unregistered, registered in nominee name.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (s *SubBalanceType13Choice) SetCode(value string) {
	s.Code = (*SecuritiesBalanceType12Code)(&value)
}

func (s *SubBalanceType13Choice) AddProprietary() *GenericIdentification47 {
	s.Proprietary = new(GenericIdentification47)
	return s.Proprietary
}
