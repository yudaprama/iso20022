package model

// Choice of format for the repair reason.
type SecuritiesTransactionType12Choice struct {

	// Securities transaction type expressed as an ISO 20022 code.
	Code *SecuritiesTransactionType10Code `xml:"Cd"`

	// Securities transaction type expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (s *SecuritiesTransactionType12Choice) SetCode(value string) {
	s.Code = (*SecuritiesTransactionType10Code)(&value)
}

func (s *SecuritiesTransactionType12Choice) AddProprietary() *GenericIdentification20 {
	s.Proprietary = new(GenericIdentification20)
	return s.Proprietary
}
