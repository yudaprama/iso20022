package model

// Choice of response status.
type ResponseStatus8Choice struct {

	// Reason provided for the status.
	Consented *ConsentStatus5Choice `xml:"Cnsntd"`

	// Rejected status of an instruction, advice or request.
	Rejected *RejectionStatus27Choice `xml:"Rjctd"`

	// Pending status of an instruction, advice or request.
	Pending *PendingStatus20Choice `xml:"Pdg"`
}

func (r *ResponseStatus8Choice) AddConsented() *ConsentStatus5Choice {
	r.Consented = new(ConsentStatus5Choice)
	return r.Consented
}

func (r *ResponseStatus8Choice) AddRejected() *RejectionStatus27Choice {
	r.Rejected = new(RejectionStatus27Choice)
	return r.Rejected
}

func (r *ResponseStatus8Choice) AddPending() *PendingStatus20Choice {
	r.Pending = new(PendingStatus20Choice)
	return r.Pending
}
