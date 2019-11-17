package model

// Choice of format for the status of payment of a security at a particular time.
type SecuritiesPaymentStatus5Choice struct {

	// Securities payment status expressed as an ISO 20022 code.
	Code *SecuritiesPaymentStatus1Code `xml:"Cd"`

	// Securities payment status expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (s *SecuritiesPaymentStatus5Choice) SetCode(value string) {
	s.Code = (*SecuritiesPaymentStatus1Code)(&value)
}

func (s *SecuritiesPaymentStatus5Choice) AddProprietary() *GenericIdentification30 {
	s.Proprietary = new(GenericIdentification30)
	return s.Proprietary
}
