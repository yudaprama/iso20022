package model

// Choice of response status.
type ResponseStatus6Choice struct {

	// Reason provided for the status.
	Consented *ConsentStatus4Choice `xml:"Cnsntd"`

	// Rejected status of an instruction, advice or request.
	Rejected *RejectionStatus20Choice `xml:"Rjctd"`

	// Pending status of an instruction, advice or request.
	Pending *PendingStatus20Choice `xml:"Pdg"`
}

func (r *ResponseStatus6Choice) AddConsented() *ConsentStatus4Choice {
	r.Consented = new(ConsentStatus4Choice)
	return r.Consented
}

func (r *ResponseStatus6Choice) AddRejected() *RejectionStatus20Choice {
	r.Rejected = new(RejectionStatus20Choice)
	return r.Rejected
}

func (r *ResponseStatus6Choice) AddPending() *PendingStatus20Choice {
	r.Pending = new(PendingStatus20Choice)
	return r.Pending
}
