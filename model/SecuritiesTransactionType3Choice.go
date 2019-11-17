package model

// Choice of format for the repair reason.
type SecuritiesTransactionType3Choice struct {

	// Securities transaction type expressed as an ISO 20022 code.
	Code *SecuritiesTransactionType4Code `xml:"Cd"`

	// Securities transaction type expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (s *SecuritiesTransactionType3Choice) SetCode(value string) {
	s.Code = (*SecuritiesTransactionType4Code)(&value)
}

func (s *SecuritiesTransactionType3Choice) AddProprietary() *GenericIdentification20 {
	s.Proprietary = new(GenericIdentification20)
	return s.Proprietary
}
