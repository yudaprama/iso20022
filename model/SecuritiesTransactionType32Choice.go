package model

// Choice of format for the settlement transaction type information.
type SecuritiesTransactionType32Choice struct {

	// Securities transaction type expressed as an ISO 20022 code.
	Code *SecuritiesTransactionType17Code `xml:"Cd"`

	// Securities transaction type expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (s *SecuritiesTransactionType32Choice) SetCode(value string) {
	s.Code = (*SecuritiesTransactionType17Code)(&value)
}

func (s *SecuritiesTransactionType32Choice) AddProprietary() *GenericIdentification30 {
	s.Proprietary = new(GenericIdentification30)
	return s.Proprietary
}
