package model

// Choice of format for the sub-balance.
type SubBalanceType9Choice struct {

	// Reason a financial instrument is not available or additional information about the financial instrument for which the balance is given, for example, unregistered, registered in nominee name.
	Code *SecuritiesBalanceType14Code `xml:"Cd"`

	// Reason a financial instrument is not available or additional information about the financial instrument for which the balance is given, for example, unregistered, registered in nominee name.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (s *SubBalanceType9Choice) SetCode(value string) {
	s.Code = (*SecuritiesBalanceType14Code)(&value)
}

func (s *SubBalanceType9Choice) AddProprietary() *GenericIdentification30 {
	s.Proprietary = new(GenericIdentification30)
	return s.Proprietary
}
