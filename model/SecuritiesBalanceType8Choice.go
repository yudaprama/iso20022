package model

// Choice of format for the balance type.
type SecuritiesBalanceType8Choice struct {

	// Sub-balance expressed as an ISO 20022 code.
	Code *SecuritiesBalanceType11Code `xml:"Cd"`

	// Sub-balance expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (s *SecuritiesBalanceType8Choice) SetCode(value string) {
	s.Code = (*SecuritiesBalanceType11Code)(&value)
}

func (s *SecuritiesBalanceType8Choice) AddProprietary() *GenericIdentification47 {
	s.Proprietary = new(GenericIdentification47)
	return s.Proprietary
}
