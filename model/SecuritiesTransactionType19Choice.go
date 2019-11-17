package model

// Choice of format for the repair reason.
type SecuritiesTransactionType19Choice struct {

	// Securities transaction type expressed as an ISO 20022 code.
	Code *SecuritiesTransactionType16Code `xml:"Cd"`

	// Securities transaction type expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (s *SecuritiesTransactionType19Choice) SetCode(value string) {
	s.Code = (*SecuritiesTransactionType16Code)(&value)
}

func (s *SecuritiesTransactionType19Choice) AddProprietary() *GenericIdentification30 {
	s.Proprietary = new(GenericIdentification30)
	return s.Proprietary
}
