package model

// Choice of reason.
type Reason16Choice struct {

	// Specifies additional information on the repurchase agreement call request acknowledgement.
	RepoCallAcknowledgementReason *AcknowledgementReason13Choice `xml:"RepoCallAckRsn,omitempty"`

	// Specifies the reason why the related instruction is cancelled, or the related cancellation request is executed.
	CancellationReason *CancellationReason20Choice `xml:"CxlRsn,omitempty"`

	// Specifies the reason why the cancellation request is pending.
	PendingCancellationReason *PendingCancellationReasons4Choice `xml:"PdgCxlRsn,omitempty"`

	// Specifies the reason why the transaction was generated.
	GeneratedReason *GeneratedReasons5Choice `xml:"GnrtdRsn,omitempty"`

	// Specifies the reason why the request was denied.
	DeniedReason *DeniedReason14Choice `xml:"DndRsn,omitempty"`

	// Specifies additional information about the processed instruction.
	AcknowledgedAcceptedReason *AcknowledgementReason12Choice `xml:"AckdAccptdRsn,omitempty"`

	// Specifies the reason why the instruction has a pending status.
	PendingReason *PendingReason29Choice `xml:"PdgRsn,omitempty"`

	// Specifies the reason why the instruction has a failing settlement status.
	FailingReason *FailingReason9Choice `xml:"FlngRsn,omitempty"`

	// Specifies the reason why the instruction has a pending processing status.
	PendingProcessingReason *PendingProcessingReason11Choice `xml:"PdgPrcgRsn,omitempty"`

	// Specifies the reason why the instruction/request has a rejected status.
	RejectionReason *RejectionReason22Choice `xml:"RjctnRsn,omitempty"`

	// Specifies the reason why the instruction is in repair.
	RepairReason *RepairReason11Choice `xml:"RprRsn,omitempty"`

	// Specifies the reason why the modification request is pending.
	PendingModificationReason *PendingReason28Choice `xml:"PdgModRsn,omitempty"`

	// Specifies the reason why the instruction has an unmatched status.
	UnmatchedReason *UnmatchedReason22Choice `xml:"UmtchdRsn,omitempty"`
}

func (r *Reason16Choice) AddRepoCallAcknowledgementReason() *AcknowledgementReason13Choice {
	r.RepoCallAcknowledgementReason = new(AcknowledgementReason13Choice)
	return r.RepoCallAcknowledgementReason
}

func (r *Reason16Choice) AddCancellationReason() *CancellationReason20Choice {
	r.CancellationReason = new(CancellationReason20Choice)
	return r.CancellationReason
}

func (r *Reason16Choice) AddPendingCancellationReason() *PendingCancellationReasons4Choice {
	r.PendingCancellationReason = new(PendingCancellationReasons4Choice)
	return r.PendingCancellationReason
}

func (r *Reason16Choice) AddGeneratedReason() *GeneratedReasons5Choice {
	r.GeneratedReason = new(GeneratedReasons5Choice)
	return r.GeneratedReason
}

func (r *Reason16Choice) AddDeniedReason() *DeniedReason14Choice {
	r.DeniedReason = new(DeniedReason14Choice)
	return r.DeniedReason
}

func (r *Reason16Choice) AddAcknowledgedAcceptedReason() *AcknowledgementReason12Choice {
	r.AcknowledgedAcceptedReason = new(AcknowledgementReason12Choice)
	return r.AcknowledgedAcceptedReason
}

func (r *Reason16Choice) AddPendingReason() *PendingReason29Choice {
	r.PendingReason = new(PendingReason29Choice)
	return r.PendingReason
}

func (r *Reason16Choice) AddFailingReason() *FailingReason9Choice {
	r.FailingReason = new(FailingReason9Choice)
	return r.FailingReason
}

func (r *Reason16Choice) AddPendingProcessingReason() *PendingProcessingReason11Choice {
	r.PendingProcessingReason = new(PendingProcessingReason11Choice)
	return r.PendingProcessingReason
}

func (r *Reason16Choice) AddRejectionReason() *RejectionReason22Choice {
	r.RejectionReason = new(RejectionReason22Choice)
	return r.RejectionReason
}

func (r *Reason16Choice) AddRepairReason() *RepairReason11Choice {
	r.RepairReason = new(RepairReason11Choice)
	return r.RepairReason
}

func (r *Reason16Choice) AddPendingModificationReason() *PendingReason28Choice {
	r.PendingModificationReason = new(PendingReason28Choice)
	return r.PendingModificationReason
}

func (r *Reason16Choice) AddUnmatchedReason() *UnmatchedReason22Choice {
	r.UnmatchedReason = new(UnmatchedReason22Choice)
	return r.UnmatchedReason
}
