package model

// Choice of format for the settlement transaction type information.
type SettlementTransactionType1Choice struct {

	// Securities transaction type expressed as an ISO 20022 code.
	Code *SettlementTransactionType7Code `xml:"Cd"`

	// Securities transaction type expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (s *SettlementTransactionType1Choice) SetCode(value string) {
	s.Code = (*SettlementTransactionType7Code)(&value)
}

func (s *SettlementTransactionType1Choice) AddProprietary() *GenericIdentification38 {
	s.Proprietary = new(GenericIdentification38)
	return s.Proprietary
}
