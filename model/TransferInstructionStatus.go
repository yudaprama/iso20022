package model

// Instruction status and the reason for the status.
type TransferInstructionStatus struct {

	// Status of the transfer instruction.
	Status *TransferStatus1Code `xml:"Sts"`

	// Additional information about the status in textual form.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (t *TransferInstructionStatus) SetStatus(value string) {
	t.Status = (*TransferStatus1Code)(&value)
}

func (t *TransferInstructionStatus) SetAdditionalInformation(value string) {
	t.AdditionalInformation = (*Max350Text)(&value)
}
