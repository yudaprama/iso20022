package model

// Instruction status.
type TransferInstructionStatus4 struct {

	// Status code.
	Status *TransferStatus4Code `xml:"Sts"`

	// Reason for the status.
	Reason *Max350Text `xml:"Rsn,omitempty"`
}

func (t *TransferInstructionStatus4) SetStatus(value string) {
	t.Status = (*TransferStatus4Code)(&value)
}

func (t *TransferInstructionStatus4) SetReason(value string) {
	t.Reason = (*Max350Text)(&value)
}
