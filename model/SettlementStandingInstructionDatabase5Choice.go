package model

// Choice of format for the standing settlement instruction information.
type SettlementStandingInstructionDatabase5Choice struct {

	// Settlement standing instruction database expressed as an ISO 20022 code.
	Code *SettlementStandingInstructionDatabase1Code `xml:"Cd"`

	// Settlement standing instruction database expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (s *SettlementStandingInstructionDatabase5Choice) SetCode(value string) {
	s.Code = (*SettlementStandingInstructionDatabase1Code)(&value)
}

func (s *SettlementStandingInstructionDatabase5Choice) AddProprietary() *GenericIdentification47 {
	s.Proprietary = new(GenericIdentification47)
	return s.Proprietary
}
