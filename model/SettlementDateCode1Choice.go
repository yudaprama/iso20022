package model

// Choice of format for the settlement date code.
type SettlementDateCode1Choice struct {

	// Settlement date expressed as an ISO 20022 code.
	Code *SettlementDate4Code `xml:"Cd"`

	// Settlement date expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (s *SettlementDateCode1Choice) SetCode(value string) {
	s.Code = (*SettlementDate4Code)(&value)
}

func (s *SettlementDateCode1Choice) AddProprietary() *GenericIdentification20 {
	s.Proprietary = new(GenericIdentification20)
	return s.Proprietary
}
