package model

// Choice of format for the settlement transaction conditions.
type SettlementTransactionCondition26Choice struct {

	// Settlement conditions expressed as an ISO 20022 code.
	Code *SettlementTransactionCondition4Code `xml:"Cd"`

	// Settlement conditions expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (s *SettlementTransactionCondition26Choice) SetCode(value string) {
	s.Code = (*SettlementTransactionCondition4Code)(&value)
}

func (s *SettlementTransactionCondition26Choice) AddProprietary() *GenericIdentification47 {
	s.Proprietary = new(GenericIdentification47)
	return s.Proprietary
}
