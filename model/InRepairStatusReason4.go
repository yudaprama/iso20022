package model

// Reason for an in repair status.
type InRepairStatusReason4 struct {

	// Reason for the in repair status expressed as a code.
	Reason *InRepairStatusReason5Choice `xml:"Rsn"`

	// Additional information about the in repair reason.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (i *InRepairStatusReason4) AddReason() *InRepairStatusReason5Choice {
	i.Reason = new(InRepairStatusReason5Choice)
	return i.Reason
}

func (i *InRepairStatusReason4) SetAdditionalInformation(value string) {
	i.AdditionalInformation = (*Max350Text)(&value)
}
