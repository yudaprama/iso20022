package model

// Choice of format for the settlement transaction conditions.
type SettlementTransactionCondition30Choice struct {

	// Settlement condition expressed as a code.
	Code *SettlementTransactionCondition11Code `xml:"Cd"`

	// Settlement condition expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (s *SettlementTransactionCondition30Choice) SetCode(value string) {
	s.Code = (*SettlementTransactionCondition11Code)(&value)
}

func (s *SettlementTransactionCondition30Choice) AddProprietary() *GenericIdentification30 {
	s.Proprietary = new(GenericIdentification30)
	return s.Proprietary
}
