package model

// Choice of format for the settlement status.
type SettlementStatus25Choice struct {

	// Provides the status of settlement of an instruction/financial instrument movement.
	Code *SecuritiesSettlementStatus2Code `xml:"Cd"`

	// Provides the status of settlement of an instruction/financial instrument movement.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (s *SettlementStatus25Choice) SetCode(value string) {
	s.Code = (*SecuritiesSettlementStatus2Code)(&value)
}

func (s *SettlementStatus25Choice) AddProprietary() *GenericIdentification47 {
	s.Proprietary = new(GenericIdentification47)
	return s.Proprietary
}
