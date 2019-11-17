package model

// Choice of format for the settlement transaction conditions.
type SettlementTransactionCondition29Choice struct {

	// Settlement conditions expressed as an ISO 20022 code.
	Code *SettlementTransactionCondition8Code `xml:"Cd"`

	// Settlement conditions expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (s *SettlementTransactionCondition29Choice) SetCode(value string) {
	s.Code = (*SettlementTransactionCondition8Code)(&value)
}

func (s *SettlementTransactionCondition29Choice) AddProprietary() *GenericIdentification47 {
	s.Proprietary = new(GenericIdentification47)
	return s.Proprietary
}
