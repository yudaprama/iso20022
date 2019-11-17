package model

// Choice of format for the processing status.
type ProcessingStatus54Choice struct {

	// Cancellation is pending processing.
	PendingCancellation *PendingStatus39Choice `xml:"PdgCxl"`

	// Instruction/Request has been rejected for further processing.
	Rejected *RejectionOrRepairStatus30Choice `xml:"Rjctd"`

	// Instruction/Request is accepted but in repair.
	Repair *RejectionOrRepairStatus31Choice `xml:"Rpr"`

	// Instruction has been acknowledged by the account servicer.
	AcknowledgedAccepted *AcknowledgedAcceptedStatus24Choice `xml:"AckdAccptd"`

	// Status that cannot be reported using one of the available standard status.
	Proprietary *ProprietaryStatusAndReason6 `xml:"Prtry"`

	// Instruction/Request will not be executed.
	Denied *DeniedStatus15Choice `xml:"Dnd"`

	// Cancellation requested executed.
	Cancelled *CancellationStatus15Choice `xml:"Canc"`
}

func (p *ProcessingStatus54Choice) AddPendingCancellation() *PendingStatus39Choice {
	p.PendingCancellation = new(PendingStatus39Choice)
	return p.PendingCancellation
}

func (p *ProcessingStatus54Choice) AddRejected() *RejectionOrRepairStatus30Choice {
	p.Rejected = new(RejectionOrRepairStatus30Choice)
	return p.Rejected
}

func (p *ProcessingStatus54Choice) AddRepair() *RejectionOrRepairStatus31Choice {
	p.Repair = new(RejectionOrRepairStatus31Choice)
	return p.Repair
}

func (p *ProcessingStatus54Choice) AddAcknowledgedAccepted() *AcknowledgedAcceptedStatus24Choice {
	p.AcknowledgedAccepted = new(AcknowledgedAcceptedStatus24Choice)
	return p.AcknowledgedAccepted
}

func (p *ProcessingStatus54Choice) AddProprietary() *ProprietaryStatusAndReason6 {
	p.Proprietary = new(ProprietaryStatusAndReason6)
	return p.Proprietary
}

func (p *ProcessingStatus54Choice) AddDenied() *DeniedStatus15Choice {
	p.Denied = new(DeniedStatus15Choice)
	return p.Denied
}

func (p *ProcessingStatus54Choice) AddCancelled() *CancellationStatus15Choice {
	p.Cancelled = new(CancellationStatus15Choice)
	return p.Cancelled
}
