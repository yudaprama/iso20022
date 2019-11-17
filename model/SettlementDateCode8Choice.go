package model

// Choice of format for the settlement date code.
type SettlementDateCode8Choice struct {

	// Settlement date expressed as an ISO 20022 code.
	Code *DateType4Code `xml:"Cd"`

	// Settlement date expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (s *SettlementDateCode8Choice) SetCode(value string) {
	s.Code = (*DateType4Code)(&value)
}

func (s *SettlementDateCode8Choice) AddProprietary() *GenericIdentification30 {
	s.Proprietary = new(GenericIdentification30)
	return s.Proprietary
}
