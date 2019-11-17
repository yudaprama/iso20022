package model

// Choice of format for the settlement transaction conditions.
type SettlementTransactionCondition17Choice struct {

	// Settlement conditions expressed as an ISO 20022 code.
	Code *SettlementTransactionCondition4Code `xml:"Cd"`

	// Settlement conditions expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (s *SettlementTransactionCondition17Choice) SetCode(value string) {
	s.Code = (*SettlementTransactionCondition4Code)(&value)
}

func (s *SettlementTransactionCondition17Choice) AddProprietary() *GenericIdentification30 {
	s.Proprietary = new(GenericIdentification30)
	return s.Proprietary
}
