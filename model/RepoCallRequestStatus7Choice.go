package model

// Choice of format for the repurchase agreement call acknowledgement.
type RepoCallRequestStatus7Choice struct {

	// Instruction has been acknowledged by the account servicer.
	AcknowledgedAccepted *AcknowledgedAcceptedStatus22Choice `xml:"AckdAccptd"`

	// Instruction/Request will not be executed.
	Denied *DeniedStatus17Choice `xml:"Dnd"`

	// Proprietary status.
	Proprietary *ProprietaryStatusAndReason6 `xml:"Prtry"`
}

func (r *RepoCallRequestStatus7Choice) AddAcknowledgedAccepted() *AcknowledgedAcceptedStatus22Choice {
	r.AcknowledgedAccepted = new(AcknowledgedAcceptedStatus22Choice)
	return r.AcknowledgedAccepted
}

func (r *RepoCallRequestStatus7Choice) AddDenied() *DeniedStatus17Choice {
	r.Denied = new(DeniedStatus17Choice)
	return r.Denied
}

func (r *RepoCallRequestStatus7Choice) AddProprietary() *ProprietaryStatusAndReason6 {
	r.Proprietary = new(ProprietaryStatusAndReason6)
	return r.Proprietary
}
