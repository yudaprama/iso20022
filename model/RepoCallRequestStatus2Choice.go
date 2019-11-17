package model

// Choice of format for the repurchase agreement call acknowledgement.
type RepoCallRequestStatus2Choice struct {

	// Instruction has been acknowledged by the account servicer.
	AcknowledgedAccepted *AcknowledgedAcceptedStatus2Choice `xml:"AckdAccptd"`

	// Instruction/Request will not be executed.
	Denied *DeniedStatus1Choice `xml:"Dnd"`

	// Proprietary status.
	Proprietary *ProprietaryStatusAndReason1 `xml:"Prtry"`
}

func (r *RepoCallRequestStatus2Choice) AddAcknowledgedAccepted() *AcknowledgedAcceptedStatus2Choice {
	r.AcknowledgedAccepted = new(AcknowledgedAcceptedStatus2Choice)
	return r.AcknowledgedAccepted
}

func (r *RepoCallRequestStatus2Choice) AddDenied() *DeniedStatus1Choice {
	r.Denied = new(DeniedStatus1Choice)
	return r.Denied
}

func (r *RepoCallRequestStatus2Choice) AddProprietary() *ProprietaryStatusAndReason1 {
	r.Proprietary = new(ProprietaryStatusAndReason1)
	return r.Proprietary
}
