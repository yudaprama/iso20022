package model

// Further information on the reversal reason of the transaction.
type ReversalReasonInformation1 struct {

	// Party issuing the reversal.
	ReversalOriginator *PartyIdentification8 `xml:"RvslOrgtr,omitempty"`

	// Specifies the reason for the reversal.
	ReversalReason *ReversalReason1Choice `xml:"RvslRsn,omitempty"`

	// Further details on the reversal reason.
	AdditionalReversalReasonInformation []*Max105Text `xml:"AddtlRvslRsnInf,omitempty"`
}

func (r *ReversalReasonInformation1) AddReversalOriginator() *PartyIdentification8 {
	r.ReversalOriginator = new(PartyIdentification8)
	return r.ReversalOriginator
}

func (r *ReversalReasonInformation1) AddReversalReason() *ReversalReason1Choice {
	r.ReversalReason = new(ReversalReason1Choice)
	return r.ReversalReason
}

func (r *ReversalReasonInformation1) AddAdditionalReversalReasonInformation(value string) {
	r.AdditionalReversalReasonInformation = append(r.AdditionalReversalReasonInformation, (*Max105Text)(&value))
}
