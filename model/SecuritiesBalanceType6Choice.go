package model

// Choice of format for the balance type.
type SecuritiesBalanceType6Choice struct {

	// Sub-balance expressed as an ISO 20022 code.
	Code *SecuritiesBalanceType11Code `xml:"Cd"`

	// Sub-balance expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (s *SecuritiesBalanceType6Choice) SetCode(value string) {
	s.Code = (*SecuritiesBalanceType11Code)(&value)
}

func (s *SecuritiesBalanceType6Choice) AddProprietary() *GenericIdentification30 {
	s.Proprietary = new(GenericIdentification30)
	return s.Proprietary
}
