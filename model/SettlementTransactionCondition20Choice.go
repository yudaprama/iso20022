package model

// Choice of format for the settlement transaction conditions.
type SettlementTransactionCondition20Choice struct {

	// Settlement conditions expressed as an ISO 20022 code.
	Code *SettlementTransactionCondition8Code `xml:"Cd"`

	// Settlement conditions expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (s *SettlementTransactionCondition20Choice) SetCode(value string) {
	s.Code = (*SettlementTransactionCondition8Code)(&value)
}

func (s *SettlementTransactionCondition20Choice) AddProprietary() *GenericIdentification30 {
	s.Proprietary = new(GenericIdentification30)
	return s.Proprietary
}
