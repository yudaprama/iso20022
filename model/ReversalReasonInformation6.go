package model

// Set of elements used to provide information on the reason of the reversal of the transaction.
type ReversalReasonInformation6 struct {

	// Party that issues the reversal.
	Originator *PartyIdentification32 `xml:"Orgtr,omitempty"`

	// Specifies the reason for the reversal.
	Reason *ReversalReason4Choice `xml:"Rsn,omitempty"`

	// Further details on the reversal reason.
	AdditionalInformation []*Max105Text `xml:"AddtlInf,omitempty"`
}

func (r *ReversalReasonInformation6) AddOriginator() *PartyIdentification32 {
	r.Originator = new(PartyIdentification32)
	return r.Originator
}

func (r *ReversalReasonInformation6) AddReason() *ReversalReason4Choice {
	r.Reason = new(ReversalReason4Choice)
	return r.Reason
}

func (r *ReversalReasonInformation6) AddAdditionalInformation(value string) {
	r.AdditionalInformation = append(r.AdditionalInformation, (*Max105Text)(&value))
}
