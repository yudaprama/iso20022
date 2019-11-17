package model

// Choice of format for the settlement date code.
type SettlementDateCode2Choice struct {

	// Settlement date expressed as an ISO 20022 code.
	Code *DateType4Code `xml:"Cd"`

	// Settlement date expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (s *SettlementDateCode2Choice) SetCode(value string) {
	s.Code = (*DateType4Code)(&value)
}

func (s *SettlementDateCode2Choice) AddProprietary() *GenericIdentification20 {
	s.Proprietary = new(GenericIdentification20)
	return s.Proprietary
}
