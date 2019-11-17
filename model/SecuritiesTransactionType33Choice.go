package model

// Choice of format for the repair reason.
type SecuritiesTransactionType33Choice struct {

	// Securities transaction type expressed as an ISO 20022 code.
	Code *SecuritiesTransactionType18Code `xml:"Cd"`

	// Securities transaction type expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (s *SecuritiesTransactionType33Choice) SetCode(value string) {
	s.Code = (*SecuritiesTransactionType18Code)(&value)
}

func (s *SecuritiesTransactionType33Choice) AddProprietary() *GenericIdentification30 {
	s.Proprietary = new(GenericIdentification30)
	return s.Proprietary
}
