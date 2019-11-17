package model

// Information about a switch order.
type SwitchOrderInstruction1 struct {

	// Information related to the switch order
	SwitchOrderDetails *SwitchOrder2 `xml:"SwtchOrdrDtls"`

	// Confirmation of the information related to an intermediary.
	IntermediaryDetails []*Intermediary4 `xml:"IntrmyDtls,omitempty"`

	// Information provided when the message is a copy of a previous message.
	CopyDetails *CopyInformation1 `xml:"CpyDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (s *SwitchOrderInstruction1) AddSwitchOrderDetails() *SwitchOrder2 {
	s.SwitchOrderDetails = new(SwitchOrder2)
	return s.SwitchOrderDetails
}

func (s *SwitchOrderInstruction1) AddIntermediaryDetails() *Intermediary4 {
	newValue := new(Intermediary4)
	s.IntermediaryDetails = append(s.IntermediaryDetails, newValue)
	return newValue
}

func (s *SwitchOrderInstruction1) AddCopyDetails() *CopyInformation1 {
	s.CopyDetails = new(CopyInformation1)
	return s.CopyDetails
}

func (s *SwitchOrderInstruction1) AddExtension() *Extension1 {
	newValue := new(Extension1)
	s.Extension = append(s.Extension, newValue)
	return newValue
}
