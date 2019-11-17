package model

// Choice of format for the standing settlement instruction information.
type SettlementStandingInstructionDatabase1Choice struct {

	// Settlement standing instruction database expressed as an ISO 20022 code.
	Code *SettlementStandingInstructionDatabase1Code `xml:"Cd"`

	// Settlement standing instruction database expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (s *SettlementStandingInstructionDatabase1Choice) SetCode(value string) {
	s.Code = (*SettlementStandingInstructionDatabase1Code)(&value)
}

func (s *SettlementStandingInstructionDatabase1Choice) AddProprietary() *GenericIdentification20 {
	s.Proprietary = new(GenericIdentification20)
	return s.Proprietary
}
