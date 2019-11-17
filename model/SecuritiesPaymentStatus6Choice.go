package model

// Choice of format for the status of payment of a security at a particular time.
type SecuritiesPaymentStatus6Choice struct {

	// Securities payment status expressed as an ISO 20022 code.
	Code *SecuritiesPaymentStatus1Code `xml:"Cd"`

	// Securities payment status expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (s *SecuritiesPaymentStatus6Choice) SetCode(value string) {
	s.Code = (*SecuritiesPaymentStatus1Code)(&value)
}

func (s *SecuritiesPaymentStatus6Choice) AddProprietary() *GenericIdentification47 {
	s.Proprietary = new(GenericIdentification47)
	return s.Proprietary
}
