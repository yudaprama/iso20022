package model

// Choice of format for the settlement transaction type information.
type SecuritiesTransactionType34Choice struct {

	// Securities transaction type expressed as an ISO 20022 code.
	Code *SecuritiesTransactionType17Code `xml:"Cd"`

	// Securities transaction type expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (s *SecuritiesTransactionType34Choice) SetCode(value string) {
	s.Code = (*SecuritiesTransactionType17Code)(&value)
}

func (s *SecuritiesTransactionType34Choice) AddProprietary() *GenericIdentification47 {
	s.Proprietary = new(GenericIdentification47)
	return s.Proprietary
}
