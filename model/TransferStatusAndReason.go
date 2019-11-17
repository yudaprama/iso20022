package model

// Information about the status of a transfer instruction and its reason.
type TransferStatusAndReason struct {

	// Business reference of the transfer instruction.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Instruction status and the reason for the status.
	Status *TransferInstructionStatus `xml:"Sts"`

	// Status of the transfer instruction is pending settlement.
	PendingSettlement *PendingSettlementStatusChoice `xml:"PdgSttlm"`

	// Status of the transfer instruction is unmatched.
	Unmatched *TransferUnmatchedStatus `xml:"Umtchd"`

	// Status is in repair.
	InRepair *InRepairStatus2Choice `xml:"InRpr"`

	// Status of the transfer instructed is rejected.
	Rejected *RejectedStatus3Choice `xml:"Rjctd"`

	// Party that initiates the status.
	StatusInitiator *PartyIdentification1Choice `xml:"StsInitr,omitempty"`
}

func (t *TransferStatusAndReason) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *TransferStatusAndReason) AddStatus() *TransferInstructionStatus {
	t.Status = new(TransferInstructionStatus)
	return t.Status
}

func (t *TransferStatusAndReason) AddPendingSettlement() *PendingSettlementStatusChoice {
	t.PendingSettlement = new(PendingSettlementStatusChoice)
	return t.PendingSettlement
}

func (t *TransferStatusAndReason) AddUnmatched() *TransferUnmatchedStatus {
	t.Unmatched = new(TransferUnmatchedStatus)
	return t.Unmatched
}

func (t *TransferStatusAndReason) AddInRepair() *InRepairStatus2Choice {
	t.InRepair = new(InRepairStatus2Choice)
	return t.InRepair
}

func (t *TransferStatusAndReason) AddRejected() *RejectedStatus3Choice {
	t.Rejected = new(RejectedStatus3Choice)
	return t.Rejected
}

func (t *TransferStatusAndReason) AddStatusInitiator() *PartyIdentification1Choice {
	t.StatusInitiator = new(PartyIdentification1Choice)
	return t.StatusInitiator
}
