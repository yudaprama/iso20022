package model

// Choice of format for the settlement transaction conditions.
type SettlementTransactionCondition11Choice struct {

	// Settlement conditions expressed as an ISO 20022 code.
	Code *SettlementTransactionCondition7Code `xml:"Cd"`

	// Settlement conditions expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (s *SettlementTransactionCondition11Choice) SetCode(value string) {
	s.Code = (*SettlementTransactionCondition7Code)(&value)
}

func (s *SettlementTransactionCondition11Choice) AddProprietary() *GenericIdentification38 {
	s.Proprietary = new(GenericIdentification38)
	return s.Proprietary
}
