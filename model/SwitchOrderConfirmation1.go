package model

// Switch order confirmation details.
type SwitchOrderConfirmation1 struct {

	// Indicates whether a confirmation amendment message will follow the confirmation cancellation instruction or not.
	AmendmentIndicator *YesNoIndicator `xml:"AmdmntInd"`

	// Information related to a switch execution.
	SwitchExecutionDetails []*SwitchExecution4 `xml:"SwtchExctnDtls"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*Extension1 `xml:"Xtnsn,omitempty"`
}

func (s *SwitchOrderConfirmation1) SetAmendmentIndicator(value string) {
	s.AmendmentIndicator = (*YesNoIndicator)(&value)
}

func (s *SwitchOrderConfirmation1) AddSwitchExecutionDetails() *SwitchExecution4 {
	newValue := new(SwitchExecution4)
	s.SwitchExecutionDetails = append(s.SwitchExecutionDetails, newValue)
	return newValue
}

func (s *SwitchOrderConfirmation1) AddExtension() *Extension1 {
	newValue := new(Extension1)
	s.Extension = append(s.Extension, newValue)
	return newValue
}
