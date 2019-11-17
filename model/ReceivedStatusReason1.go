package model

// Reason for a received status.
type ReceivedStatusReason1 struct {

	// Reason for the received status.
	Reason *ReceivedReason1Choice `xml:"Rsn"`

	// Additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *ReceivedStatusReason1) AddReason() *ReceivedReason1Choice {
	r.Reason = new(ReceivedReason1Choice)
	return r.Reason
}

func (r *ReceivedStatusReason1) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
