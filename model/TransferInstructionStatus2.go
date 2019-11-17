package model

// Instruction status.
type TransferInstructionStatus2 struct {

	// Status of the transfer is accepted, sent to next party, matched, already executed, or settled.
	Status *TransferStatus2Code `xml:"Sts"`

	// Reason for the status.
	Reason *Max350Text `xml:"Rsn,omitempty"`
}

func (t *TransferInstructionStatus2) SetStatus(value string) {
	t.Status = (*TransferStatus2Code)(&value)
}

func (t *TransferInstructionStatus2) SetReason(value string) {
	t.Reason = (*Max350Text)(&value)
}
