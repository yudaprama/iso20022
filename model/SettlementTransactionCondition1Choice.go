package model

// Choice of format for the settlement transaction conditions.
type SettlementTransactionCondition1Choice struct {

	// Settlement conditions expressed as an ISO 20022 code.
	Code *SettlementTransactionCondition2Code `xml:"Cd"`

	// Settlement conditions expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (s *SettlementTransactionCondition1Choice) SetCode(value string) {
	s.Code = (*SettlementTransactionCondition2Code)(&value)
}

func (s *SettlementTransactionCondition1Choice) AddProprietary() *GenericIdentification20 {
	s.Proprietary = new(GenericIdentification20)
	return s.Proprietary
}
