package model

// Choice of format for the repair reason.
type SecuritiesTransactionType24Choice struct {

	// Securities transaction type expressed as an ISO 20022 code.
	Code *SecuritiesTransactionType8Code `xml:"Cd"`

	// Securities transaction type expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (s *SecuritiesTransactionType24Choice) SetCode(value string) {
	s.Code = (*SecuritiesTransactionType8Code)(&value)
}

func (s *SecuritiesTransactionType24Choice) AddProprietary() *GenericIdentification47 {
	s.Proprietary = new(GenericIdentification47)
	return s.Proprietary
}
