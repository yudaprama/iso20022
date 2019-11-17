package model

// Choice of format for the repair reason.
type SecuritiesTransactionType30Choice struct {

	// Securities transaction type expressed as an ISO 20022 code.
	Code *SecuritiesTransactionType16Code `xml:"Cd"`

	// Securities transaction type expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (s *SecuritiesTransactionType30Choice) SetCode(value string) {
	s.Code = (*SecuritiesTransactionType16Code)(&value)
}

func (s *SecuritiesTransactionType30Choice) AddProprietary() *GenericIdentification47 {
	s.Proprietary = new(GenericIdentification47)
	return s.Proprietary
}
