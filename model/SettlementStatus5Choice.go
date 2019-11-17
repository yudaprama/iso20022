package model

// Choice of format for the settlement status.
type SettlementStatus5Choice struct {

	// Provides the status of settlement of an instruction/financial instrument movement.
	Code *SecuritiesSettlementStatus2Code `xml:"Cd"`

	// Provides the status of settlement of an instruction/financial instrument movement.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (s *SettlementStatus5Choice) SetCode(value string) {
	s.Code = (*SecuritiesSettlementStatus2Code)(&value)
}

func (s *SettlementStatus5Choice) AddProprietary() *GenericIdentification20 {
	s.Proprietary = new(GenericIdentification20)
	return s.Proprietary
}
