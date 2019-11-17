package model

// Reason for a status in repair.
type InRepairStatusReason1 struct {

	// Reason for an in repair status in free format text.
	AdditionalInformation *Max350Text `xml:"AddtlInf"`
}

func (i *InRepairStatusReason1) SetAdditionalInformation(value string) {
	i.AdditionalInformation = (*Max350Text)(&value)
}
