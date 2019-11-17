package model

// Choice of format for the settlement transaction conditions.
type SettlementTransactionCondition28Choice struct {

	// Settlement conditions expressed as an ISO 20022 code.
	Code *SettlementTransactionCondition10Code `xml:"Cd"`

	// Settlement conditions expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (s *SettlementTransactionCondition28Choice) SetCode(value string) {
	s.Code = (*SettlementTransactionCondition10Code)(&value)
}

func (s *SettlementTransactionCondition28Choice) AddProprietary() *GenericIdentification47 {
	s.Proprietary = new(GenericIdentification47)
	return s.Proprietary
}
