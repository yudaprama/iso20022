package model

// Further information on the return reason of the transaction.
type ReturnReasonInformation1 struct {

	// Party issuing the return.
	ReturnOriginator *PartyIdentification8 `xml:"RtrOrgtr,omitempty"`

	// Specifies the reason for the return.
	ReturnReason *ReturnReason1Choice `xml:"RtrRsn,omitempty"`

	// Further details on the return reason.
	AdditionalReturnReasonInformation []*Max105Text `xml:"AddtlRtrRsnInf,omitempty"`
}

func (r *ReturnReasonInformation1) AddReturnOriginator() *PartyIdentification8 {
	r.ReturnOriginator = new(PartyIdentification8)
	return r.ReturnOriginator
}

func (r *ReturnReasonInformation1) AddReturnReason() *ReturnReason1Choice {
	r.ReturnReason = new(ReturnReason1Choice)
	return r.ReturnReason
}

func (r *ReturnReasonInformation1) AddAdditionalReturnReasonInformation(value string) {
	r.AdditionalReturnReasonInformation = append(r.AdditionalReturnReasonInformation, (*Max105Text)(&value))
}
