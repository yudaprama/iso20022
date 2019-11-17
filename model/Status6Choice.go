package model

// Specifies the status.
type Status6Choice struct {

	// Status expressed as a code.
	Code *TradeStatus4Code `xml:"Cd"`

	// Status expressed as a proprietary form.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (s *Status6Choice) SetCode(value string) {
	s.Code = (*TradeStatus4Code)(&value)
}

func (s *Status6Choice) SetProprietary(value string) {
	s.Proprietary = (*Max35Text)(&value)
}
