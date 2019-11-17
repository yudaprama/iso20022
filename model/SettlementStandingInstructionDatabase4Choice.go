package model

// Choice of format for the standing settlement instruction information.
type SettlementStandingInstructionDatabase4Choice struct {

	// Settlement standing instruction database expressed as an ISO 20022 code.
	Code *SettlementStandingInstructionDatabase1Code `xml:"Cd"`

	// Settlement standing instruction database expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (s *SettlementStandingInstructionDatabase4Choice) SetCode(value string) {
	s.Code = (*SettlementStandingInstructionDatabase1Code)(&value)
}

func (s *SettlementStandingInstructionDatabase4Choice) AddProprietary() *GenericIdentification30 {
	s.Proprietary = new(GenericIdentification30)
	return s.Proprietary
}
