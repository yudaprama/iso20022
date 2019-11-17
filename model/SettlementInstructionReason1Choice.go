package model

// Choice of formats for the settlement instruction reason.
type SettlementInstructionReason1Choice struct {

	// Settlement instruction reason expressed as a code.
	Code *SettlementInstructionReason1Code `xml:"Cd"`

	// Settlement instruction reason expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (s *SettlementInstructionReason1Choice) SetCode(value string) {
	s.Code = (*SettlementInstructionReason1Code)(&value)
}

func (s *SettlementInstructionReason1Choice) AddProprietary() *GenericIdentification47 {
	s.Proprietary = new(GenericIdentification47)
	return s.Proprietary
}
