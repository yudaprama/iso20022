package model

// Choice of format for the settlement date code.
type SettlementDateCode5Choice struct {

	// Settlement date expressed as an ISO 20022 code.
	Code *SettlementDate5Code `xml:"Cd"`

	// Settlement date expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (s *SettlementDateCode5Choice) SetCode(value string) {
	s.Code = (*SettlementDate5Code)(&value)
}

func (s *SettlementDateCode5Choice) AddProprietary() *GenericIdentification38 {
	s.Proprietary = new(GenericIdentification38)
	return s.Proprietary
}
