package model

// Choice of format for the settlement date code.
type SettlementDateCode11Choice struct {

	// Settlement date expressed as an ISO 20022 code.
	Code *DateType4Code `xml:"Cd"`

	// Settlement date expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (s *SettlementDateCode11Choice) SetCode(value string) {
	s.Code = (*DateType4Code)(&value)
}

func (s *SettlementDateCode11Choice) AddProprietary() *GenericIdentification47 {
	s.Proprietary = new(GenericIdentification47)
	return s.Proprietary
}
