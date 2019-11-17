package model

// Choice of format for the settlement transaction conditions.
type SettlementTransactionCondition5Choice struct {

	// Settlement conditions expressed as an ISO 20022 code.
	Code *SettlementTransactionCondition4Code `xml:"Cd"`

	// Settlement conditions expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (s *SettlementTransactionCondition5Choice) SetCode(value string) {
	s.Code = (*SettlementTransactionCondition4Code)(&value)
}

func (s *SettlementTransactionCondition5Choice) AddProprietary() *GenericIdentification20 {
	s.Proprietary = new(GenericIdentification20)
	return s.Proprietary
}
