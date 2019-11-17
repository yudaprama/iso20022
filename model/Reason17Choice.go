package model

// Choice of reason.
type Reason17Choice struct {

	// Specifies additional information on the repurchase agreement call request acknowledgement.
	RepoCallAcknowledgementReason *AcknowledgementReason18Choice `xml:"RepoCallAckRsn,omitempty"`

	// Specifies the reason why the related instruction is cancelled, or the related cancellation request is executed.
	CancellationReason *CancellationReason30Choice `xml:"CxlRsn,omitempty"`

	// Specifies the reason why the cancellation request is pending.
	PendingCancellationReason *PendingCancellationReasons5Choice `xml:"PdgCxlRsn,omitempty"`

	// Specifies the reason why the transaction was generated.
	GeneratedReason *GeneratedReasons6Choice `xml:"GnrtdRsn,omitempty"`

	// Specifies the reason why the request was denied.
	DeniedReason *DeniedReason23Choice `xml:"DndRsn,omitempty"`

	// Specifies additional information about the processed instruction.
	AcknowledgedAcceptedReason *AcknowledgementReason16Choice `xml:"AckdAccptdRsn,omitempty"`

	// Specifies the reason why the instruction has a pending status.
	PendingReason *PendingReason47Choice `xml:"PdgRsn,omitempty"`

	// Specifies the reason why the instruction has a failing settlement status.
	FailingReason *FailingReason15Choice `xml:"FlngRsn,omitempty"`

	// Specifies the reason why the instruction has a pending processing status.
	PendingProcessingReason *PendingProcessingReason13Choice `xml:"PdgPrcgRsn,omitempty"`

	// Specifies the reason why the instruction/request has a rejected status.
	RejectionReason *RejectionReason34Choice `xml:"RjctnRsn,omitempty"`

	// Specifies the reason why the instruction is in repair.
	RepairReason *RepairReason18Choice `xml:"RprRsn,omitempty"`

	// Specifies the reason why the modification request is pending.
	PendingModificationReason *PendingReason37Choice `xml:"PdgModRsn,omitempty"`

	// Specifies the reason why the instruction has an unmatched status.
	UnmatchedReason *UnmatchedReason29Choice `xml:"UmtchdRsn,omitempty"`
}

func (r *Reason17Choice) AddRepoCallAcknowledgementReason() *AcknowledgementReason18Choice {
	r.RepoCallAcknowledgementReason = new(AcknowledgementReason18Choice)
	return r.RepoCallAcknowledgementReason
}

func (r *Reason17Choice) AddCancellationReason() *CancellationReason30Choice {
	r.CancellationReason = new(CancellationReason30Choice)
	return r.CancellationReason
}

func (r *Reason17Choice) AddPendingCancellationReason() *PendingCancellationReasons5Choice {
	r.PendingCancellationReason = new(PendingCancellationReasons5Choice)
	return r.PendingCancellationReason
}

func (r *Reason17Choice) AddGeneratedReason() *GeneratedReasons6Choice {
	r.GeneratedReason = new(GeneratedReasons6Choice)
	return r.GeneratedReason
}

func (r *Reason17Choice) AddDeniedReason() *DeniedReason23Choice {
	r.DeniedReason = new(DeniedReason23Choice)
	return r.DeniedReason
}

func (r *Reason17Choice) AddAcknowledgedAcceptedReason() *AcknowledgementReason16Choice {
	r.AcknowledgedAcceptedReason = new(AcknowledgementReason16Choice)
	return r.AcknowledgedAcceptedReason
}

func (r *Reason17Choice) AddPendingReason() *PendingReason47Choice {
	r.PendingReason = new(PendingReason47Choice)
	return r.PendingReason
}

func (r *Reason17Choice) AddFailingReason() *FailingReason15Choice {
	r.FailingReason = new(FailingReason15Choice)
	return r.FailingReason
}

func (r *Reason17Choice) AddPendingProcessingReason() *PendingProcessingReason13Choice {
	r.PendingProcessingReason = new(PendingProcessingReason13Choice)
	return r.PendingProcessingReason
}

func (r *Reason17Choice) AddRejectionReason() *RejectionReason34Choice {
	r.RejectionReason = new(RejectionReason34Choice)
	return r.RejectionReason
}

func (r *Reason17Choice) AddRepairReason() *RepairReason18Choice {
	r.RepairReason = new(RepairReason18Choice)
	return r.RepairReason
}

func (r *Reason17Choice) AddPendingModificationReason() *PendingReason37Choice {
	r.PendingModificationReason = new(PendingReason37Choice)
	return r.PendingModificationReason
}

func (r *Reason17Choice) AddUnmatchedReason() *UnmatchedReason29Choice {
	r.UnmatchedReason = new(UnmatchedReason29Choice)
	return r.UnmatchedReason
}
