package model

// Choice of response status.
type ResponseStatus3Choice struct {

	// Reason provided for the status.
	Consented *ConsentStatus2Choice `xml:"Cnsntd"`

	// The status of an instruction, advice or request.
	Rejected *RejectionStatus7Choice `xml:"Rjctd"`

	// The status of an instruction, advice or request.
	Pending *PendingStatus20Choice `xml:"Pdg"`
}

func (r *ResponseStatus3Choice) AddConsented() *ConsentStatus2Choice {
	r.Consented = new(ConsentStatus2Choice)
	return r.Consented
}

func (r *ResponseStatus3Choice) AddRejected() *RejectionStatus7Choice {
	r.Rejected = new(RejectionStatus7Choice)
	return r.Rejected
}

func (r *ResponseStatus3Choice) AddPending() *PendingStatus20Choice {
	r.Pending = new(PendingStatus20Choice)
	return r.Pending
}
