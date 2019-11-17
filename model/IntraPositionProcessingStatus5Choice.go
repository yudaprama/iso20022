package model

// Choice of format for the processing status.
type IntraPositionProcessingStatus5Choice struct {

	// Instruction/Request has been rejected for further processing.
	Rejected *RejectionOrRepairStatus29Choice `xml:"Rjctd"`

	// Instruction/Request is accepted but in repair.
	Repair *RejectionOrRepairStatus29Choice `xml:"Rpr"`

	// Instruction has been cancelled.
	Cancelled *CancellationStatus14Choice `xml:"Canc"`

	// Instruction has been acknowledged by the account servicer.
	AcknowledgedAccepted *AcknowledgedAcceptedStatus21Choice `xml:"AckdAccptd"`

	// Specifies a choice of status for the processing of an intra-position movement.
	Proprietary *ProprietaryStatusAndReason6 `xml:"Prtry"`
}

func (i *IntraPositionProcessingStatus5Choice) AddRejected() *RejectionOrRepairStatus29Choice {
	i.Rejected = new(RejectionOrRepairStatus29Choice)
	return i.Rejected
}

func (i *IntraPositionProcessingStatus5Choice) AddRepair() *RejectionOrRepairStatus29Choice {
	i.Repair = new(RejectionOrRepairStatus29Choice)
	return i.Repair
}

func (i *IntraPositionProcessingStatus5Choice) AddCancelled() *CancellationStatus14Choice {
	i.Cancelled = new(CancellationStatus14Choice)
	return i.Cancelled
}

func (i *IntraPositionProcessingStatus5Choice) AddAcknowledgedAccepted() *AcknowledgedAcceptedStatus21Choice {
	i.AcknowledgedAccepted = new(AcknowledgedAcceptedStatus21Choice)
	return i.AcknowledgedAccepted
}

func (i *IntraPositionProcessingStatus5Choice) AddProprietary() *ProprietaryStatusAndReason6 {
	i.Proprietary = new(ProprietaryStatusAndReason6)
	return i.Proprietary
}
