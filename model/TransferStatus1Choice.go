package model

// Choice of statuses for the status of the transfer.
type TransferStatus1Choice struct {

	// Status of the transfer is received, accepted, sent to next party, matched, already executed, or settled.
	Status *TransferInstructionStatus3 `xml:"Sts"`

	// Status of the transfer is pending settlement.
	PendingSettlement *PendingSettlementStatus2 `xml:"PdgSttlm"`

	// Status of the transfer is unmatched.
	Unmatched *TransferUnmatchedStatus2 `xml:"Umtchd"`

	// Status of the transfer is in repair.
	InRepair *InRepairStatus3 `xml:"InRpr"`

	// Status of the transfer is rejected.
	Rejected []*RejectedStatus8Choice `xml:"Rjctd"`

	// Status of the transfer is failed settlement, that is, settlement in the International Central Securities Depository (ICSD) or Central Securities Depository (CSD) failed.
	FailedSettlement *FailedSettlementStatus1 `xml:"FaildSttlm"`

	// Status of the transfer is cancelled.
	Cancelled *CancelledStatus3 `xml:"Canc"`

	// Status of the transfer is reversed.
	Reversed *ReversedStatus1 `xml:"Rvsd"`

	// Status of the transfer is cancellation pending.
	CancellationPending *CancellationPendingStatus1 `xml:"CxlPdg"`
}

func (t *TransferStatus1Choice) AddStatus() *TransferInstructionStatus3 {
	t.Status = new(TransferInstructionStatus3)
	return t.Status
}

func (t *TransferStatus1Choice) AddPendingSettlement() *PendingSettlementStatus2 {
	t.PendingSettlement = new(PendingSettlementStatus2)
	return t.PendingSettlement
}

func (t *TransferStatus1Choice) AddUnmatched() *TransferUnmatchedStatus2 {
	t.Unmatched = new(TransferUnmatchedStatus2)
	return t.Unmatched
}

func (t *TransferStatus1Choice) AddInRepair() *InRepairStatus3 {
	t.InRepair = new(InRepairStatus3)
	return t.InRepair
}

func (t *TransferStatus1Choice) AddRejected() *RejectedStatus8Choice {
	newValue := new(RejectedStatus8Choice)
	t.Rejected = append(t.Rejected, newValue)
	return newValue
}

func (t *TransferStatus1Choice) AddFailedSettlement() *FailedSettlementStatus1 {
	t.FailedSettlement = new(FailedSettlementStatus1)
	return t.FailedSettlement
}

func (t *TransferStatus1Choice) AddCancelled() *CancelledStatus3 {
	t.Cancelled = new(CancelledStatus3)
	return t.Cancelled
}

func (t *TransferStatus1Choice) AddReversed() *ReversedStatus1 {
	t.Reversed = new(ReversedStatus1)
	return t.Reversed
}

func (t *TransferStatus1Choice) AddCancellationPending() *CancellationPendingStatus1 {
	t.CancellationPending = new(CancellationPendingStatus1)
	return t.CancellationPending
}
