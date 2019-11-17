package model

// Choice of format for the standing settlement instruction information.
type SettlementStandingInstructionDatabase3Choice struct {

	// Settlement standing instruction database expressed as an ISO 20022 code.
	Code *SettlementStandingInstructionDatabase1Code `xml:"Cd"`

	// Settlement standing instruction database expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (s *SettlementStandingInstructionDatabase3Choice) SetCode(value string) {
	s.Code = (*SettlementStandingInstructionDatabase1Code)(&value)
}

func (s *SettlementStandingInstructionDatabase3Choice) AddProprietary() *GenericIdentification38 {
	s.Proprietary = new(GenericIdentification38)
	return s.Proprietary
}
