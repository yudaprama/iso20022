package model

// Choice of format for the settlement transaction type information.
type SecuritiesTransactionType21Choice struct {

	// Securities transaction type expressed as an ISO 20022 code.
	Code *SecuritiesTransactionType7Code `xml:"Cd"`

	// Securities transaction type expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (s *SecuritiesTransactionType21Choice) SetCode(value string) {
	s.Code = (*SecuritiesTransactionType7Code)(&value)
}

func (s *SecuritiesTransactionType21Choice) AddProprietary() *GenericIdentification30 {
	s.Proprietary = new(GenericIdentification30)
	return s.Proprietary
}
