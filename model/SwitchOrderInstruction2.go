package model

// Information about a switch order.
type SwitchOrderInstruction2 struct {

	// Information related to the switch order
	SwitchOrderDetails []*SwitchOrder3 `xml:"SwtchOrdrDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (s *SwitchOrderInstruction2) AddSwitchOrderDetails() *SwitchOrder3 {
	newValue := new(SwitchOrder3)
	s.SwitchOrderDetails = append(s.SwitchOrderDetails, newValue)
	return newValue
}

func (s *SwitchOrderInstruction2) AddExtension() *Extension1 {
	newValue := new(Extension1)
	s.Extension = append(s.Extension, newValue)
	return newValue
}
