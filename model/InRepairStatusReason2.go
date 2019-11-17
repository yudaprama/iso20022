package model

// Reason for an in repair status.
type InRepairStatusReason2 struct {

	// Reason for the in repair status in textual form.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (i *InRepairStatusReason2) SetAdditionalInformation(value string) {
	i.AdditionalInformation = (*Max350Text)(&value)
}
