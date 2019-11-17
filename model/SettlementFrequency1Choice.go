package model

// Choice of formats for the settlement frequency.
type SettlementFrequency1Choice struct {

	// Settlement frequency expressed as a code.
	Code *EventFrequency10Code `xml:"Cd"`

	// Settlement frequency expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (s *SettlementFrequency1Choice) SetCode(value string) {
	s.Code = (*EventFrequency10Code)(&value)
}

func (s *SettlementFrequency1Choice) AddProprietary() *GenericIdentification47 {
	s.Proprietary = new(GenericIdentification47)
	return s.Proprietary
}
