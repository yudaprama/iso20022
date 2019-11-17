package model

// Choice of formats to  express the type of securities balance.
type SecuritiesBalanceType10FormatChoice struct {

	// Standard code to specify the type of securities balance.
	Code *SecuritiesBalanceType10Code `xml:"Cd"`

	// Proprietary code to  express the type of securities balance.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (s *SecuritiesBalanceType10FormatChoice) SetCode(value string) {
	s.Code = (*SecuritiesBalanceType10Code)(&value)
}

func (s *SecuritiesBalanceType10FormatChoice) AddProprietary() *GenericIdentification13 {
	s.Proprietary = new(GenericIdentification13)
	return s.Proprietary
}
