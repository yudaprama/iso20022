package model

// Instruction status.
type TransferInstructionStatus3 struct {

	// Status code.
	Status *TransferStatus3Code `xml:"Sts"`

	// Reason for the status.
	Reason *Max350Text `xml:"Rsn,omitempty"`
}

func (t *TransferInstructionStatus3) SetStatus(value string) {
	t.Status = (*TransferStatus3Code)(&value)
}

func (t *TransferInstructionStatus3) SetReason(value string) {
	t.Reason = (*Max350Text)(&value)
}
