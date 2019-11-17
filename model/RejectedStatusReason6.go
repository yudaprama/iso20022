package model

// Reason for a rejected status.
type RejectedStatusReason6 struct {

	// Reason for a rejected status in structured form.
	Structured *RejectedStatusReason5Code `xml:"Strd"`

	// Reason for a rejected status in free format text.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (r *RejectedStatusReason6) SetStructured(value string) {
	r.Structured = (*RejectedStatusReason5Code)(&value)
}

func (r *RejectedStatusReason6) SetAdditionalInformation(value string) {
	r.AdditionalInformation = (*Max350Text)(&value)
}
