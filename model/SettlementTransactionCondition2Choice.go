package model

// Choice of format for the settlement transaction conditions.
type SettlementTransactionCondition2Choice struct {

	// Settlement conditions expressed as an ISO 20022 code.
	Code *SettlementTransactionCondition3Code `xml:"Cd"`

	// Settlement conditions expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (s *SettlementTransactionCondition2Choice) SetCode(value string) {
	s.Code = (*SettlementTransactionCondition3Code)(&value)
}

func (s *SettlementTransactionCondition2Choice) AddProprietary() *GenericIdentification20 {
	s.Proprietary = new(GenericIdentification20)
	return s.Proprietary
}
