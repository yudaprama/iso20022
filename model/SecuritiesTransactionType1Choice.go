package model

// Choice of format for the settlement transaction type information.
type SecuritiesTransactionType1Choice struct {

	// Securities transaction type expressed as an ISO 20022 code.
	Code *SecuritiesTransactionType1Code `xml:"Cd"`

	// Securities transaction type expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (s *SecuritiesTransactionType1Choice) SetCode(value string) {
	s.Code = (*SecuritiesTransactionType1Code)(&value)
}

func (s *SecuritiesTransactionType1Choice) AddProprietary() *GenericIdentification20 {
	s.Proprietary = new(GenericIdentification20)
	return s.Proprietary
}
