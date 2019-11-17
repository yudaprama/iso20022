package model

// Choice of formats for the specification of the statement frequency.
type StatementFrequencyReason2Choice struct {

	// Statement frequency expressed as a code.
	Code *EventFrequency9Code `xml:"Cd"`

	// Statement frequency expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (s *StatementFrequencyReason2Choice) SetCode(value string) {
	s.Code = (*EventFrequency9Code)(&value)
}

func (s *StatementFrequencyReason2Choice) AddProprietary() *GenericIdentification47 {
	s.Proprietary = new(GenericIdentification47)
	return s.Proprietary
}
