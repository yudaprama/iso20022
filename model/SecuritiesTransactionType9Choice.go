package model

// Choice of format for the settlement transaction type information.
type SecuritiesTransactionType9Choice struct {

	// Securities transaction type expressed as an ISO 20022 code.
	Code *SecuritiesTransactionType7Code `xml:"Cd"`

	// Securities transaction type expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (s *SecuritiesTransactionType9Choice) SetCode(value string) {
	s.Code = (*SecuritiesTransactionType7Code)(&value)
}

func (s *SecuritiesTransactionType9Choice) AddProprietary() *GenericIdentification20 {
	s.Proprietary = new(GenericIdentification20)
	return s.Proprietary
}
