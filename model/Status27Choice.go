package model

// Specifies the status.
type Status27Choice struct {

	// Status expressed as a code.
	Code *TradeStatus6Code `xml:"Cd"`

	// Status expressed in a proprietary form.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (s *Status27Choice) SetCode(value string) {
	s.Code = (*TradeStatus6Code)(&value)
}

func (s *Status27Choice) SetProprietary(value string) {
	s.Proprietary = (*Max35Text)(&value)
}
