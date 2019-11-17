package model

// Choice of format for the balance type.
type SecuritiesBalanceType11Choice struct {

	// Sub-balance expressed as an ISO 20022 code.
	Code *SecuritiesBalanceType13Code `xml:"Cd"`

	// Sub-balance expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (s *SecuritiesBalanceType11Choice) SetCode(value string) {
	s.Code = (*SecuritiesBalanceType13Code)(&value)
}

func (s *SecuritiesBalanceType11Choice) AddProprietary() *GenericIdentification47 {
	s.Proprietary = new(GenericIdentification47)
	return s.Proprietary
}
