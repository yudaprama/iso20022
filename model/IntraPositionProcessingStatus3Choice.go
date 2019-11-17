package model

// Choice of format for the processing status.
type IntraPositionProcessingStatus3Choice struct {

	// Instruction/Request has been rejected for further processing.
	Rejected *RejectionOrRepairStatus15Choice `xml:"Rjctd"`

	// Instruction/Request is accepted but in repair.
	Repair *RejectionOrRepairStatus15Choice `xml:"Rpr"`

	// Instruction has been cancelled (only as an response to an SecuritiesTransactionStatusQuery). The status on the processing of a cancellation request must be provided using a SecuritiesTransactionCancellationRequestStatusAdvice.
	Cancelled *CancellationStatus7Choice `xml:"Canc"`

	// Instruction has been acknowledged by the account servicer.
	AcknowledgedAccepted *AcknowledgedAcceptedStatus7Choice `xml:"AckdAccptd"`

	// Specifies a choice of status for the processing of an intra-position movement.
	Proprietary *ProprietaryStatusAndReason1 `xml:"Prtry"`
}

func (i *IntraPositionProcessingStatus3Choice) AddRejected() *RejectionOrRepairStatus15Choice {
	i.Rejected = new(RejectionOrRepairStatus15Choice)
	return i.Rejected
}

func (i *IntraPositionProcessingStatus3Choice) AddRepair() *RejectionOrRepairStatus15Choice {
	i.Repair = new(RejectionOrRepairStatus15Choice)
	return i.Repair
}

func (i *IntraPositionProcessingStatus3Choice) AddCancelled() *CancellationStatus7Choice {
	i.Cancelled = new(CancellationStatus7Choice)
	return i.Cancelled
}

func (i *IntraPositionProcessingStatus3Choice) AddAcknowledgedAccepted() *AcknowledgedAcceptedStatus7Choice {
	i.AcknowledgedAccepted = new(AcknowledgedAcceptedStatus7Choice)
	return i.AcknowledgedAccepted
}

func (i *IntraPositionProcessingStatus3Choice) AddProprietary() *ProprietaryStatusAndReason1 {
	i.Proprietary = new(ProprietaryStatusAndReason1)
	return i.Proprietary
}
