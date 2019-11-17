package model

// Choice of format for the repair reason.
type SecuritiesTransactionType10Choice struct {

	// Securities transaction type expressed as an ISO 20022 code.
	Code *SecuritiesTransactionType8Code `xml:"Cd"`

	// Securities transaction type expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (s *SecuritiesTransactionType10Choice) SetCode(value string) {
	s.Code = (*SecuritiesTransactionType8Code)(&value)
}

func (s *SecuritiesTransactionType10Choice) AddProprietary() *GenericIdentification20 {
	s.Proprietary = new(GenericIdentification20)
	return s.Proprietary
}
