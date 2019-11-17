package model

// Choice of format for the settlement transaction conditions.
type SettlementTransactionCondition18Choice struct {

	// Settlement conditions expressed as an ISO 20022 code.
	Code *SettlementTransactionCondition6Code `xml:"Cd"`

	// Settlement conditions expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (s *SettlementTransactionCondition18Choice) SetCode(value string) {
	s.Code = (*SettlementTransactionCondition6Code)(&value)
}

func (s *SettlementTransactionCondition18Choice) AddProprietary() *GenericIdentification30 {
	s.Proprietary = new(GenericIdentification30)
	return s.Proprietary
}
