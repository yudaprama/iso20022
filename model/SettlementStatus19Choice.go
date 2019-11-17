package model

// Choice of format for the settlement status.
type SettlementStatus19Choice struct {

	// Provides the status of settlement of an instruction/financial instrument movement.
	Code *SecuritiesSettlementStatus2Code `xml:"Cd"`

	// Provides the status of settlement of an instruction/financial instrument movement.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (s *SettlementStatus19Choice) SetCode(value string) {
	s.Code = (*SecuritiesSettlementStatus2Code)(&value)
}

func (s *SettlementStatus19Choice) AddProprietary() *GenericIdentification30 {
	s.Proprietary = new(GenericIdentification30)
	return s.Proprietary
}
