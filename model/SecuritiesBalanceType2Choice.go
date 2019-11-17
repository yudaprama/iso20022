package model

// Choice of format for the balance type.
type SecuritiesBalanceType2Choice struct {

	// Sub-balance expressed as an ISO 20022 code.
	Code *SecuritiesBalanceType13Code `xml:"Cd"`

	// Sub-balance expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (s *SecuritiesBalanceType2Choice) SetCode(value string) {
	s.Code = (*SecuritiesBalanceType13Code)(&value)
}

func (s *SecuritiesBalanceType2Choice) AddProprietary() *GenericIdentification20 {
	s.Proprietary = new(GenericIdentification20)
	return s.Proprietary
}
