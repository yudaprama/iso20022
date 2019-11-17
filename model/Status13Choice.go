package model

// Specifies the status.
type Status13Choice struct {

	// Status expressed as a code.
	Code *TradeStatus5Code `xml:"Cd"`

	// Status expressed in a proprietary form.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (s *Status13Choice) SetCode(value string) {
	s.Code = (*TradeStatus5Code)(&value)
}

func (s *Status13Choice) SetProprietary(value string) {
	s.Proprietary = (*Max35Text)(&value)
}
