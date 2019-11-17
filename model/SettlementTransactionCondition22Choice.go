package model

// Choice of format for the settlement transaction conditions.
type SettlementTransactionCondition22Choice struct {

	// Settlement conditions expressed as an ISO 20022 code.
	Code *SettlementTransactionCondition6Code `xml:"Cd"`

	// Settlement conditions expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (s *SettlementTransactionCondition22Choice) SetCode(value string) {
	s.Code = (*SettlementTransactionCondition6Code)(&value)
}

func (s *SettlementTransactionCondition22Choice) AddProprietary() *GenericIdentification47 {
	s.Proprietary = new(GenericIdentification47)
	return s.Proprietary
}
