package model

// Choice of statuses for the status of the transfer.
type TransferStatus2Choice struct {

	// Status of the transfer is received, accepted, sent to next party, matched, already executed, or settled.
	Status *TransferInstructionStatus4 `xml:"Sts"`

	// Status of the transfer is pending settlement.
	PendingSettlement *PendingSettlementStatus3Choice `xml:"PdgSttlm"`

	// Status of the transfer is unmatched.
	Unmatched *TransferUnmatchedStatus3Choice `xml:"Umtchd"`

	// Status of the transfer is in repair.
	InRepair *InRepairStatus4Choice `xml:"InRpr"`

	// Status of the transfer is rejected.
	Rejected []*RejectionReason32 `xml:"Rjctd"`

	// Status of the transfer is failed settlement, that is, settlement in the International Central Securities Depository (ICSD) or Central Securities Depository (CSD) failed.
	FailedSettlement *FailedSettlementStatus2Choice `xml:"FaildSttlm"`

	// Status of the transfer is cancelled.
	Cancelled *CancelledStatus13Choice `xml:"Canc"`

	// Status of the transfer is reversed.
	Reversed *ReversedStatus2Choice `xml:"Rvsd"`

	// Status of the transfer is cancellation pending.
	CancellationPending *CancellationPendingStatus7Choice `xml:"CxlPdg"`
}

func (t *TransferStatus2Choice) AddStatus() *TransferInstructionStatus4 {
	t.Status = new(TransferInstructionStatus4)
	return t.Status
}

func (t *TransferStatus2Choice) AddPendingSettlement() *PendingSettlementStatus3Choice {
	t.PendingSettlement = new(PendingSettlementStatus3Choice)
	return t.PendingSettlement
}

func (t *TransferStatus2Choice) AddUnmatched() *TransferUnmatchedStatus3Choice {
	t.Unmatched = new(TransferUnmatchedStatus3Choice)
	return t.Unmatched
}

func (t *TransferStatus2Choice) AddInRepair() *InRepairStatus4Choice {
	t.InRepair = new(InRepairStatus4Choice)
	return t.InRepair
}

func (t *TransferStatus2Choice) AddRejected() *RejectionReason32 {
	newValue := new(RejectionReason32)
	t.Rejected = append(t.Rejected, newValue)
	return newValue
}

func (t *TransferStatus2Choice) AddFailedSettlement() *FailedSettlementStatus2Choice {
	t.FailedSettlement = new(FailedSettlementStatus2Choice)
	return t.FailedSettlement
}

func (t *TransferStatus2Choice) AddCancelled() *CancelledStatus13Choice {
	t.Cancelled = new(CancelledStatus13Choice)
	return t.Cancelled
}

func (t *TransferStatus2Choice) AddReversed() *ReversedStatus2Choice {
	t.Reversed = new(ReversedStatus2Choice)
	return t.Reversed
}

func (t *TransferStatus2Choice) AddCancellationPending() *CancellationPendingStatus7Choice {
	t.CancellationPending = new(CancellationPendingStatus7Choice)
	return t.CancellationPending
}
