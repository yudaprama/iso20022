package model

// Instruction for a change to an account status and reason for the change.
type AccountStatusUpdateInstruction1 struct {

	// Type of status change instructed for the account status.
	UpdateInstruction *AccountStatusUpdateInstruction1Choice `xml:"UpdInstr"`

	// Reason for the instruction to change the account status.
	UpdateInstructionReason *AccountStatusUpdateInstructionReason1Choice `xml:"UpdInstrRsn,omitempty"`
}

func (a *AccountStatusUpdateInstruction1) AddUpdateInstruction() *AccountStatusUpdateInstruction1Choice {
	a.UpdateInstruction = new(AccountStatusUpdateInstruction1Choice)
	return a.UpdateInstruction
}

func (a *AccountStatusUpdateInstruction1) AddUpdateInstructionReason() *AccountStatusUpdateInstructionReason1Choice {
	a.UpdateInstructionReason = new(AccountStatusUpdateInstructionReason1Choice)
	return a.UpdateInstructionReason
}
