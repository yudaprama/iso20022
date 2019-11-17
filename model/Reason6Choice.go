package model

// Choice of reason.
type Reason6Choice struct {

	// Specifies additional information on the repurchase agreement call request acknowledgement.
	RepoCallAcknowledgementReason *AcknowledgementReason3Choice `xml:"RepoCallAckRsn,omitempty"`

	// Specifies the reason why the related instruction is cancelled, or the related cancellation request is executed.
	CancellationReason *CancellationReason9Choice `xml:"CxlRsn,omitempty"`

	// Specifies the reason why the cancellation request is pending.
	PendingCancellationReason *PendingCancellationReasons2Choice `xml:"PdgCxlRsn,omitempty"`

	// Specifies the reason why the transaction was generated.
	GeneratedReason *GeneratedReasons1Choice `xml:"GnrtdRsn,omitempty"`

	// Specifies the reason why the request was denied.
	DeniedReason *DeniedReason1Choice `xml:"DndRsn,omitempty"`

	// Specifies additional information about the processed instruction.
	AcknowledgedAcceptedReason *AcknowledgementReason2Choice `xml:"AckdAccptdRsn,omitempty"`

	// Specifies the reason why the instruction has a pending status.
	PendingReason *PendingReason11Choice `xml:"PdgRsn,omitempty"`

	// Specifies the reason why the instruction has a failing settlement status.
	FailingReason *FailingReason1Choice `xml:"FlngRsn,omitempty"`

	// Specifies the reason why the instruction has a pending processing status.
	PendingProcessingReason *PendingProcessingReason1Choice `xml:"PdgPrcgRsn,omitempty"`

	// Specifies the reason why the instruction/request has a rejected status.
	RejectionReason *RejectionReason5Choice `xml:"RjctnRsn,omitempty"`

	// Specifies the reason why the instruction is in repair.
	RepairReason *RepairReason7Choice `xml:"RprRsn,omitempty"`

	// Specifies the reason why the modification request is pending.
	PendingModificationReason *PendingReason2Choice `xml:"PdgModRsn,omitempty"`

	// Specifies the reason why the instruction has an unmatched status.
	UnmatchedReason *UnmatchedReason2Choice `xml:"UmtchdRsn,omitempty"`
}

func (r *Reason6Choice) AddRepoCallAcknowledgementReason() *AcknowledgementReason3Choice {
	r.RepoCallAcknowledgementReason = new(AcknowledgementReason3Choice)
	return r.RepoCallAcknowledgementReason
}

func (r *Reason6Choice) AddCancellationReason() *CancellationReason9Choice {
	r.CancellationReason = new(CancellationReason9Choice)
	return r.CancellationReason
}

func (r *Reason6Choice) AddPendingCancellationReason() *PendingCancellationReasons2Choice {
	r.PendingCancellationReason = new(PendingCancellationReasons2Choice)
	return r.PendingCancellationReason
}

func (r *Reason6Choice) AddGeneratedReason() *GeneratedReasons1Choice {
	r.GeneratedReason = new(GeneratedReasons1Choice)
	return r.GeneratedReason
}

func (r *Reason6Choice) AddDeniedReason() *DeniedReason1Choice {
	r.DeniedReason = new(DeniedReason1Choice)
	return r.DeniedReason
}

func (r *Reason6Choice) AddAcknowledgedAcceptedReason() *AcknowledgementReason2Choice {
	r.AcknowledgedAcceptedReason = new(AcknowledgementReason2Choice)
	return r.AcknowledgedAcceptedReason
}

func (r *Reason6Choice) AddPendingReason() *PendingReason11Choice {
	r.PendingReason = new(PendingReason11Choice)
	return r.PendingReason
}

func (r *Reason6Choice) AddFailingReason() *FailingReason1Choice {
	r.FailingReason = new(FailingReason1Choice)
	return r.FailingReason
}

func (r *Reason6Choice) AddPendingProcessingReason() *PendingProcessingReason1Choice {
	r.PendingProcessingReason = new(PendingProcessingReason1Choice)
	return r.PendingProcessingReason
}

func (r *Reason6Choice) AddRejectionReason() *RejectionReason5Choice {
	r.RejectionReason = new(RejectionReason5Choice)
	return r.RejectionReason
}

func (r *Reason6Choice) AddRepairReason() *RepairReason7Choice {
	r.RepairReason = new(RepairReason7Choice)
	return r.RepairReason
}

func (r *Reason6Choice) AddPendingModificationReason() *PendingReason2Choice {
	r.PendingModificationReason = new(PendingReason2Choice)
	return r.PendingModificationReason
}

func (r *Reason6Choice) AddUnmatchedReason() *UnmatchedReason2Choice {
	r.UnmatchedReason = new(UnmatchedReason2Choice)
	return r.UnmatchedReason
}
