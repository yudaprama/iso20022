package model

// Choice of reason.
type Reason1Choice struct {

	// Specifies additional information on the repurchase agreement call request acknowledgement.
	RepoCallAcknowledgementReason *AcknowledgementReason3Choice `xml:"RepoCallAckRsn,omitempty"`

	// Specifies the reason why the related instruction is cancelled, or the related cancellation request is executed.
	CancellationReason *CancellationReason4Choice `xml:"CxlRsn,omitempty"`

	// Specifies the reason why the cancellation request is pending.
	PendingCancellationReason *PendingCancellationReasons2Choice `xml:"PdgCxlRsn,omitempty"`

	// Specifies the reason why the transaction was generated.
	GeneratedReason *GeneratedReasons1Choice `xml:"GnrtdRsn,omitempty"`

	// Specifies the reason why the request was denied.
	DeniedReason *DeniedReason1Choice `xml:"DndRsn,omitempty"`

	// Specifies additional information about the processed instruction.
	AcknowledgedAcceptedReason *AcknowledgementReason2Choice `xml:"AckdAccptdRsn,omitempty"`

	// Specifies the reason why the instruction has a pending status.
	PendingReason *PendingReason3Choice `xml:"PdgRsn,omitempty"`

	// Specifies the reason why the instruction has a failing settlement status.
	FailingReason *FailingReason1Choice `xml:"FlngRsn,omitempty"`

	// Specifies the reason why the instruction has a pending processing status.
	PendingProcessingReason *PendingProcessingReason1Choice `xml:"PdgPrcgRsn,omitempty"`

	// Specifies the reason why the instruction/request has a rejected status.
	RejectionReason *RejectionReason5Choice `xml:"RjctnRsn,omitempty"`

	// Specifies the reason why the instruction is in repair.
	RepairReason *RepairReason3Choice `xml:"RprRsn,omitempty"`

	// Specifies the reason why the modification request is pending.
	PendingModificationReason *PendingReason2Choice `xml:"PdgModRsn,omitempty"`

	// Specifies the reason why the instruction has an unmatched status.
	UnmatchedReason *UnmatchedReason2Choice `xml:"UmtchdRsn,omitempty"`
}

func (r *Reason1Choice) AddRepoCallAcknowledgementReason() *AcknowledgementReason3Choice {
	r.RepoCallAcknowledgementReason = new(AcknowledgementReason3Choice)
	return r.RepoCallAcknowledgementReason
}

func (r *Reason1Choice) AddCancellationReason() *CancellationReason4Choice {
	r.CancellationReason = new(CancellationReason4Choice)
	return r.CancellationReason
}

func (r *Reason1Choice) AddPendingCancellationReason() *PendingCancellationReasons2Choice {
	r.PendingCancellationReason = new(PendingCancellationReasons2Choice)
	return r.PendingCancellationReason
}

func (r *Reason1Choice) AddGeneratedReason() *GeneratedReasons1Choice {
	r.GeneratedReason = new(GeneratedReasons1Choice)
	return r.GeneratedReason
}

func (r *Reason1Choice) AddDeniedReason() *DeniedReason1Choice {
	r.DeniedReason = new(DeniedReason1Choice)
	return r.DeniedReason
}

func (r *Reason1Choice) AddAcknowledgedAcceptedReason() *AcknowledgementReason2Choice {
	r.AcknowledgedAcceptedReason = new(AcknowledgementReason2Choice)
	return r.AcknowledgedAcceptedReason
}

func (r *Reason1Choice) AddPendingReason() *PendingReason3Choice {
	r.PendingReason = new(PendingReason3Choice)
	return r.PendingReason
}

func (r *Reason1Choice) AddFailingReason() *FailingReason1Choice {
	r.FailingReason = new(FailingReason1Choice)
	return r.FailingReason
}

func (r *Reason1Choice) AddPendingProcessingReason() *PendingProcessingReason1Choice {
	r.PendingProcessingReason = new(PendingProcessingReason1Choice)
	return r.PendingProcessingReason
}

func (r *Reason1Choice) AddRejectionReason() *RejectionReason5Choice {
	r.RejectionReason = new(RejectionReason5Choice)
	return r.RejectionReason
}

func (r *Reason1Choice) AddRepairReason() *RepairReason3Choice {
	r.RepairReason = new(RepairReason3Choice)
	return r.RepairReason
}

func (r *Reason1Choice) AddPendingModificationReason() *PendingReason2Choice {
	r.PendingModificationReason = new(PendingReason2Choice)
	return r.PendingModificationReason
}

func (r *Reason1Choice) AddUnmatchedReason() *UnmatchedReason2Choice {
	r.UnmatchedReason = new(UnmatchedReason2Choice)
	return r.UnmatchedReason
}
