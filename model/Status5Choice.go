package model

// Specifies the status.
type Status5Choice struct {

	// Status expressed as a code.
	Code *TradeStatus3Code `xml:"Cd"`

	// Status expressed as a proprietary form.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (s *Status5Choice) SetCode(value string) {
	s.Code = (*TradeStatus3Code)(&value)
}

func (s *Status5Choice) SetProprietary(value string) {
	s.Proprietary = (*Max35Text)(&value)
}
