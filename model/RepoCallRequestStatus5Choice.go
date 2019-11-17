package model

// Choice of format for the repurchase agreement call acknowledgement.
type RepoCallRequestStatus5Choice struct {

	// Instruction has been acknowledged by the account servicer.
	AcknowledgedAccepted *AcknowledgedAcceptedStatus10Choice `xml:"AckdAccptd"`

	// Instruction/Request will not be executed.
	Denied *DeniedStatus5Choice `xml:"Dnd"`

	// Proprietary status.
	Proprietary *ProprietaryStatusAndReason1 `xml:"Prtry"`
}

func (r *RepoCallRequestStatus5Choice) AddAcknowledgedAccepted() *AcknowledgedAcceptedStatus10Choice {
	r.AcknowledgedAccepted = new(AcknowledgedAcceptedStatus10Choice)
	return r.AcknowledgedAccepted
}

func (r *RepoCallRequestStatus5Choice) AddDenied() *DeniedStatus5Choice {
	r.Denied = new(DeniedStatus5Choice)
	return r.Denied
}

func (r *RepoCallRequestStatus5Choice) AddProprietary() *ProprietaryStatusAndReason1 {
	r.Proprietary = new(ProprietaryStatusAndReason1)
	return r.Proprietary
}
