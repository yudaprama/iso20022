package model

// Reason for a rejected status.
type RejectedStatusReason5 struct {

	// Reason for a rejected status in structured form.
	Structured *TransferRejectedStatusReason1Code `xml:"Strd"`

	// Additional information about the reason for the rejected status in textual form.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (r *RejectedStatusReason5) SetStructured(value string) {
	r.Structured = (*TransferRejectedStatusReason1Code)(&value)
}

func (r *RejectedStatusReason5) SetAdditionalInformation(value string) {
	r.AdditionalInformation = (*Max350Text)(&value)
}
