package model

// Choice of format for the balance type.
type SecuritiesBalanceType7Choice struct {

	// Sub-balance expressed as an ISO 20022 code.
	Code *SecuritiesBalanceType13Code `xml:"Cd"`

	// Sub-balance expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (s *SecuritiesBalanceType7Choice) SetCode(value string) {
	s.Code = (*SecuritiesBalanceType13Code)(&value)
}

func (s *SecuritiesBalanceType7Choice) AddProprietary() *GenericIdentification30 {
	s.Proprietary = new(GenericIdentification30)
	return s.Proprietary
}
