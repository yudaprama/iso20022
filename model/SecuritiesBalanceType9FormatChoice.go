package model

// Choice of formats to  express the type of securities balance.
type SecuritiesBalanceType9FormatChoice struct {

	// Standard code to specify the type of securities balance.
	Code *SecuritiesBalanceType9Code `xml:"Cd"`

	// Proprietary code to  express the type of securities balance.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (s *SecuritiesBalanceType9FormatChoice) SetCode(value string) {
	s.Code = (*SecuritiesBalanceType9Code)(&value)
}

func (s *SecuritiesBalanceType9FormatChoice) AddProprietary() *GenericIdentification13 {
	s.Proprietary = new(GenericIdentification13)
	return s.Proprietary
}
