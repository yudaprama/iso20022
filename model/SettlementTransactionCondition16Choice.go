package model

// Choice of format for the settlement transaction conditions.
type SettlementTransactionCondition16Choice struct {

	// Settlement conditions expressed as an ISO 20022 code.
	Code *SettlementTransactionCondition10Code `xml:"Cd"`

	// Settlement conditions expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (s *SettlementTransactionCondition16Choice) SetCode(value string) {
	s.Code = (*SettlementTransactionCondition10Code)(&value)
}

func (s *SettlementTransactionCondition16Choice) AddProprietary() *GenericIdentification30 {
	s.Proprietary = new(GenericIdentification30)
	return s.Proprietary
}
