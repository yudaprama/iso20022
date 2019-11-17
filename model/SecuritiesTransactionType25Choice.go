package model

// Choice of format for the settlement transaction type information.
type SecuritiesTransactionType25Choice struct {

	// Securities transaction type expressed as an ISO 20022 code.
	Code *SecuritiesTransactionType7Code `xml:"Cd"`

	// Securities transaction type expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (s *SecuritiesTransactionType25Choice) SetCode(value string) {
	s.Code = (*SecuritiesTransactionType7Code)(&value)
}

func (s *SecuritiesTransactionType25Choice) AddProprietary() *GenericIdentification47 {
	s.Proprietary = new(GenericIdentification47)
	return s.Proprietary
}
