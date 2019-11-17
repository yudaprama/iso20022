package model

// Choice of format for the settlement date code.
type SettlementDateCode9Choice struct {

	// Settlement date expressed as an ISO 20022 code.
	Code *SettlementDate4Code `xml:"Cd"`

	// Settlement date expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (s *SettlementDateCode9Choice) SetCode(value string) {
	s.Code = (*SettlementDate4Code)(&value)
}

func (s *SettlementDateCode9Choice) AddProprietary() *GenericIdentification47 {
	s.Proprietary = new(GenericIdentification47)
	return s.Proprietary
}
