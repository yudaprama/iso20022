package model

// Choice of formats to  express the type of securities balance.
type SecuritiesBalanceType6FormatChoice struct {

	// Standard code to specify the type of securities balance.
	Code *SecuritiesBalanceType6Code `xml:"Cd"`

	// Proprietary code to  express the type of securities balance.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (s *SecuritiesBalanceType6FormatChoice) SetCode(value string) {
	s.Code = (*SecuritiesBalanceType6Code)(&value)
}

func (s *SecuritiesBalanceType6FormatChoice) AddProprietary() *GenericIdentification13 {
	s.Proprietary = new(GenericIdentification13)
	return s.Proprietary
}
