package model

// Choice of format for the settlement instruction generation.
type SettlementInstructionGeneration1Choice struct {

	// Settlement instruction generation expressed as a ISO20022 code.
	Code *SettlementInstructionGeneration1Code `xml:"Cd"`

	// Settlement instruction generation expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (s *SettlementInstructionGeneration1Choice) SetCode(value string) {
	s.Code = (*SettlementInstructionGeneration1Code)(&value)
}

func (s *SettlementInstructionGeneration1Choice) AddProprietary() *GenericIdentification38 {
	s.Proprietary = new(GenericIdentification38)
	return s.Proprietary
}
