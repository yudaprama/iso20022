package model

// Choice of format for the processing status.
type IntraPositionProcessingStatus1Choice struct {

	// Instruction/Request has been rejected for further processing.
	Rejected *RejectionOrRepairStatus2Choice `xml:"Rjctd"`

	// Instruction/Request is accepted but in repair.
	Repair *RejectionOrRepairStatus2Choice `xml:"Rpr"`

	// Instruction has been cancelled (only as an response to an SecuritiesTransactionStatusQuery). The status on the processing of a cancellation request must be provided using a SecuritiesTransactionCancellationRequestStatusAdvice.
	Cancelled *CancellationStatus4Choice `xml:"Canc"`

	// Instruction has been acknowledged by the account servicer.
	AcknowledgedAccepted *AcknowledgedAcceptedStatus3Choice `xml:"AckdAccptd"`

	// Specifies a choice of status for the processing of an intra-position movement.
	Proprietary *ProprietaryStatusAndReason1 `xml:"Prtry"`
}

func (i *IntraPositionProcessingStatus1Choice) AddRejected() *RejectionOrRepairStatus2Choice {
	i.Rejected = new(RejectionOrRepairStatus2Choice)
	return i.Rejected
}

func (i *IntraPositionProcessingStatus1Choice) AddRepair() *RejectionOrRepairStatus2Choice {
	i.Repair = new(RejectionOrRepairStatus2Choice)
	return i.Repair
}

func (i *IntraPositionProcessingStatus1Choice) AddCancelled() *CancellationStatus4Choice {
	i.Cancelled = new(CancellationStatus4Choice)
	return i.Cancelled
}

func (i *IntraPositionProcessingStatus1Choice) AddAcknowledgedAccepted() *AcknowledgedAcceptedStatus3Choice {
	i.AcknowledgedAccepted = new(AcknowledgedAcceptedStatus3Choice)
	return i.AcknowledgedAccepted
}

func (i *IntraPositionProcessingStatus1Choice) AddProprietary() *ProprietaryStatusAndReason1 {
	i.Proprietary = new(ProprietaryStatusAndReason1)
	return i.Proprietary
}
