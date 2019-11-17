package model

// Reason for a rejected status.
type RejectedStatusReason12 struct {

	// Reason for the rejected status.
	Reason *RejectedReason8Choice `xml:"Rsn"`

	// Additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RejectedStatusReason12) AddReason() *RejectedReason8Choice {
	r.Reason = new(RejectedReason8Choice)
	return r.Reason
}

func (r *RejectedStatusReason12) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
