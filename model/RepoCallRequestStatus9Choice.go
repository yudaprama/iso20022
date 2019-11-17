package model

// Choice of format for the repurchase agreement call acknowledgement.
type RepoCallRequestStatus9Choice struct {

	// Instruction has been acknowledged by the account servicer.
	AcknowledgedAccepted *AcknowledgedAcceptedStatus27Choice `xml:"AckdAccptd"`

	// Instruction/Request will not be executed.
	Denied *DeniedStatus18Choice `xml:"Dnd"`

	// Proprietary status.
	Proprietary *ProprietaryStatusAndReason7 `xml:"Prtry"`
}

func (r *RepoCallRequestStatus9Choice) AddAcknowledgedAccepted() *AcknowledgedAcceptedStatus27Choice {
	r.AcknowledgedAccepted = new(AcknowledgedAcceptedStatus27Choice)
	return r.AcknowledgedAccepted
}

func (r *RepoCallRequestStatus9Choice) AddDenied() *DeniedStatus18Choice {
	r.Denied = new(DeniedStatus18Choice)
	return r.Denied
}

func (r *RepoCallRequestStatus9Choice) AddProprietary() *ProprietaryStatusAndReason7 {
	r.Proprietary = new(ProprietaryStatusAndReason7)
	return r.Proprietary
}
