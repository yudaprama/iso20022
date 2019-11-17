package model

// Choice of format for the processing status.
type IntraPositionProcessingStatus6Choice struct {

	// Instruction/Request has been rejected for further processing.
	Rejected *RejectionOrRepairStatus33Choice `xml:"Rjctd"`

	// Instruction/Request is accepted but in repair.
	Repair *RejectionOrRepairStatus33Choice `xml:"Rpr"`

	// Instruction has been cancelled.
	Cancelled *CancellationStatus17Choice `xml:"Canc"`

	// Instruction has been acknowledged by the account servicer.
	AcknowledgedAccepted *AcknowledgedAcceptedStatus25Choice `xml:"AckdAccptd"`

	// Specifies a choice of status for the processing of an intra-position movement.
	Proprietary *ProprietaryStatusAndReason7 `xml:"Prtry"`
}

func (i *IntraPositionProcessingStatus6Choice) AddRejected() *RejectionOrRepairStatus33Choice {
	i.Rejected = new(RejectionOrRepairStatus33Choice)
	return i.Rejected
}

func (i *IntraPositionProcessingStatus6Choice) AddRepair() *RejectionOrRepairStatus33Choice {
	i.Repair = new(RejectionOrRepairStatus33Choice)
	return i.Repair
}

func (i *IntraPositionProcessingStatus6Choice) AddCancelled() *CancellationStatus17Choice {
	i.Cancelled = new(CancellationStatus17Choice)
	return i.Cancelled
}

func (i *IntraPositionProcessingStatus6Choice) AddAcknowledgedAccepted() *AcknowledgedAcceptedStatus25Choice {
	i.AcknowledgedAccepted = new(AcknowledgedAcceptedStatus25Choice)
	return i.AcknowledgedAccepted
}

func (i *IntraPositionProcessingStatus6Choice) AddProprietary() *ProprietaryStatusAndReason7 {
	i.Proprietary = new(ProprietaryStatusAndReason7)
	return i.Proprietary
}
