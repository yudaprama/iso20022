package model

// Choice of format for the balance type.
type SecuritiesBalanceType3Choice struct {

	// Sub-balance expressed as an ISO 20022 code.
	Code *SecuritiesBalanceType11Code `xml:"Cd"`

	// Sub-balance expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (s *SecuritiesBalanceType3Choice) SetCode(value string) {
	s.Code = (*SecuritiesBalanceType11Code)(&value)
}

func (s *SecuritiesBalanceType3Choice) AddProprietary() *GenericIdentification20 {
	s.Proprietary = new(GenericIdentification20)
	return s.Proprietary
}
