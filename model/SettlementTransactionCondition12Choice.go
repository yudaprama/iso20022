package model

// Choice of format for the settlement transaction conditions.
type SettlementTransactionCondition12Choice struct {

	// Settlement conditions expressed as an ISO 20022 code.
	Code *SettlementTransactionCondition8Code `xml:"Cd"`

	// Settlement conditions expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (s *SettlementTransactionCondition12Choice) SetCode(value string) {
	s.Code = (*SettlementTransactionCondition8Code)(&value)
}

func (s *SettlementTransactionCondition12Choice) AddProprietary() *GenericIdentification20 {
	s.Proprietary = new(GenericIdentification20)
	return s.Proprietary
}
