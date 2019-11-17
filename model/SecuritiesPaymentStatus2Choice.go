package model

// Choice of format for the status of payment of a security at a particular time.
type SecuritiesPaymentStatus2Choice struct {

	// Securities payment status expressed as an ISO 20022 code.
	Code *SecuritiesPaymentStatus1Code `xml:"Cd"`

	// Securities payment status expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (s *SecuritiesPaymentStatus2Choice) SetCode(value string) {
	s.Code = (*SecuritiesPaymentStatus1Code)(&value)
}

func (s *SecuritiesPaymentStatus2Choice) AddProprietary() *GenericIdentification20 {
	s.Proprietary = new(GenericIdentification20)
	return s.Proprietary
}
