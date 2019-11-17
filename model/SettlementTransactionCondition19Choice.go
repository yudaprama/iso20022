package model

// Choice of format for the settlement transaction conditions.
type SettlementTransactionCondition19Choice struct {

	// Settlement conditions expressed as an ISO 20022 code.
	Code *SettlementTransactionCondition3Code `xml:"Cd"`

	// Settlement conditions expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (s *SettlementTransactionCondition19Choice) SetCode(value string) {
	s.Code = (*SettlementTransactionCondition3Code)(&value)
}

func (s *SettlementTransactionCondition19Choice) AddProprietary() *GenericIdentification30 {
	s.Proprietary = new(GenericIdentification30)
	return s.Proprietary
}
