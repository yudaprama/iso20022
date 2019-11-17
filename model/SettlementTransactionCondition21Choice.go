package model

// Choice of format for the settlement transaction conditions.
type SettlementTransactionCondition21Choice struct {

	// Settlement conditions expressed as an ISO 20022 code.
	Code *SettlementTransactionCondition3Code `xml:"Cd"`

	// Settlement conditions expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (s *SettlementTransactionCondition21Choice) SetCode(value string) {
	s.Code = (*SettlementTransactionCondition3Code)(&value)
}

func (s *SettlementTransactionCondition21Choice) AddProprietary() *GenericIdentification47 {
	s.Proprietary = new(GenericIdentification47)
	return s.Proprietary
}
