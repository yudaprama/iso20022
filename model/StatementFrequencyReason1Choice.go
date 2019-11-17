package model

// Choice of formats for the specification of the statement frequency.
type StatementFrequencyReason1Choice struct {

	// Statement frequency expressed as a code.
	Code *EventFrequency1Code `xml:"Cd"`

	// Statement frequency expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (s *StatementFrequencyReason1Choice) SetCode(value string) {
	s.Code = (*EventFrequency1Code)(&value)
}

func (s *StatementFrequencyReason1Choice) AddProprietary() *GenericIdentification47 {
	s.Proprietary = new(GenericIdentification47)
	return s.Proprietary
}
