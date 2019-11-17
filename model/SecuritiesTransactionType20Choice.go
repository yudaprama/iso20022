package model

// Choice of format for the repair reason.
type SecuritiesTransactionType20Choice struct {

	// Securities transaction type expressed as an ISO 20022 code.
	Code *SecuritiesTransactionType10Code `xml:"Cd"`

	// Securities transaction type expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (s *SecuritiesTransactionType20Choice) SetCode(value string) {
	s.Code = (*SecuritiesTransactionType10Code)(&value)
}

func (s *SecuritiesTransactionType20Choice) AddProprietary() *GenericIdentification30 {
	s.Proprietary = new(GenericIdentification30)
	return s.Proprietary
}
