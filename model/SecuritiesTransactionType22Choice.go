package model

// Choice of format for the repair reason.
type SecuritiesTransactionType22Choice struct {

	// Securities transaction type expressed as an ISO 20022 code.
	Code *SecuritiesTransactionType10Code `xml:"Cd"`

	// Securities transaction type expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (s *SecuritiesTransactionType22Choice) SetCode(value string) {
	s.Code = (*SecuritiesTransactionType10Code)(&value)
}

func (s *SecuritiesTransactionType22Choice) AddProprietary() *GenericIdentification47 {
	s.Proprietary = new(GenericIdentification47)
	return s.Proprietary
}
