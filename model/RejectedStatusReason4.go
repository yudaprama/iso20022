package model

// Reason for a rejected status.
type RejectedStatusReason4 struct {

	// Reason for a rejected status in structured form.
	Structured *RejectedStatusReason4Code `xml:"Strd"`

	// Reason for a rejected status in free format text.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (r *RejectedStatusReason4) SetStructured(value string) {
	r.Structured = (*RejectedStatusReason4Code)(&value)
}

func (r *RejectedStatusReason4) SetAdditionalInformation(value string) {
	r.AdditionalInformation = (*Max350Text)(&value)
}
