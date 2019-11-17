package model

// Choice of format for the settlement transaction conditions.
type SettlementTransactionCondition7Choice struct {

	// Settlement conditions expressed as an ISO 20022 code.
	Code *SettlementTransactionCondition6Code `xml:"Cd"`

	// Settlement conditions expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (s *SettlementTransactionCondition7Choice) SetCode(value string) {
	s.Code = (*SettlementTransactionCondition6Code)(&value)
}

func (s *SettlementTransactionCondition7Choice) AddProprietary() *GenericIdentification20 {
	s.Proprietary = new(GenericIdentification20)
	return s.Proprietary
}
