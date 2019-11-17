package model

// Choice of format for the repair reason.
type SecuritiesTransactionType18Choice struct {

	// Securities transaction type expressed as an ISO 20022 code.
	Code *SecuritiesTransactionType8Code `xml:"Cd"`

	// Securities transaction type expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (s *SecuritiesTransactionType18Choice) SetCode(value string) {
	s.Code = (*SecuritiesTransactionType8Code)(&value)
}

func (s *SecuritiesTransactionType18Choice) AddProprietary() *GenericIdentification30 {
	s.Proprietary = new(GenericIdentification30)
	return s.Proprietary
}
