package model

// Choice of format for the repair reason.
type SecuritiesTransactionType7Choice struct {

	// Securities transaction type expressed as an ISO 20022 code.
	Code *SecuritiesTransactionType6Code `xml:"Cd"`

	// Securities transaction type expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (s *SecuritiesTransactionType7Choice) SetCode(value string) {
	s.Code = (*SecuritiesTransactionType6Code)(&value)
}

func (s *SecuritiesTransactionType7Choice) AddProprietary() *GenericIdentification20 {
	s.Proprietary = new(GenericIdentification20)
	return s.Proprietary
}
