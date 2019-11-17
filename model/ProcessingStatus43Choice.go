package model

// Choice of formats for the specification of a status.
type ProcessingStatus43Choice struct {

	// Status of the standing settlement instruction, cancellation or deletion is received for further processing.
	Received *ReceivedStatusReason1 `xml:"Rcvd"`

	// Status of the standing settlement instruction, cancellation or deletion is
	// acknowledged/accepted for further processing. The instruction has been received, is processable and has been validated for further processing.
	Accepted *AcceptedStatusReason7 `xml:"Accptd"`

	// Status of the standing settlement instruction, cancellation or deletion is pending.
	PendingProcessing *PendingProcessingStatusReason1 `xml:"PdgPrcg"`

	// Status of the standing settlement instruction, cancellation or deletion is rejected.
	Rejected *RejectedStatusReason12 `xml:"Rjctd"`

	// Status of the standing settlement instruction, cancellation or deletion is expressed in a bilaterally agreed manner.
	ProprietaryStatus *ProprietaryStatusAndReason5 `xml:"PrtrySts"`
}

func (p *ProcessingStatus43Choice) AddReceived() *ReceivedStatusReason1 {
	p.Received = new(ReceivedStatusReason1)
	return p.Received
}

func (p *ProcessingStatus43Choice) AddAccepted() *AcceptedStatusReason7 {
	p.Accepted = new(AcceptedStatusReason7)
	return p.Accepted
}

func (p *ProcessingStatus43Choice) AddPendingProcessing() *PendingProcessingStatusReason1 {
	p.PendingProcessing = new(PendingProcessingStatusReason1)
	return p.PendingProcessing
}

func (p *ProcessingStatus43Choice) AddRejected() *RejectedStatusReason12 {
	p.Rejected = new(RejectedStatusReason12)
	return p.Rejected
}

func (p *ProcessingStatus43Choice) AddProprietaryStatus() *ProprietaryStatusAndReason5 {
	p.ProprietaryStatus = new(ProprietaryStatusAndReason5)
	return p.ProprietaryStatus
}
