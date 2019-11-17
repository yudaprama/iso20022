package model

// Choice of format for the repair reason.
type SecuritiesTransactionType11Choice struct {

	// Securities transaction type expressed as an ISO 20022 code.
	Code *SecuritiesTransactionType9Code `xml:"Cd"`

	// Securities transaction type expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (s *SecuritiesTransactionType11Choice) SetCode(value string) {
	s.Code = (*SecuritiesTransactionType9Code)(&value)
}

func (s *SecuritiesTransactionType11Choice) AddProprietary() *GenericIdentification20 {
	s.Proprietary = new(GenericIdentification20)
	return s.Proprietary
}
