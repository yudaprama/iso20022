package model

// Specifies the status.
type Status28Choice struct {

	// Status expressed as a code.
	Code *TradeStatus7Code `xml:"Cd"`

	// Status expressed as a proprietary form.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (s *Status28Choice) SetCode(value string) {
	s.Code = (*TradeStatus7Code)(&value)
}

func (s *Status28Choice) SetProprietary(value string) {
	s.Proprietary = (*Max35Text)(&value)
}
