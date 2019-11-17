package model

// Choice of format for the repair reason.
type SecuritiesTransactionType2Choice struct {

	// Securities transaction type expressed as an ISO 20022 code.
	Code *SecuritiesTransactionType3Code `xml:"Cd"`

	// Securities transaction type expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (s *SecuritiesTransactionType2Choice) SetCode(value string) {
	s.Code = (*SecuritiesTransactionType3Code)(&value)
}

func (s *SecuritiesTransactionType2Choice) AddProprietary() *GenericIdentification20 {
	s.Proprietary = new(GenericIdentification20)
	return s.Proprietary
}
