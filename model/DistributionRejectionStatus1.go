package model

// Provides information about the rejection status of a global movement.
type DistributionRejectionStatus1 struct {

	// The rejection reason.
	Reason []*RejectionReason19FormatChoice `xml:"Rsn"`

	// Additional information about the status.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (d *DistributionRejectionStatus1) AddReason() *RejectionReason19FormatChoice {
	newValue := new(RejectionReason19FormatChoice)
	d.Reason = append(d.Reason, newValue)
	return newValue
}

func (d *DistributionRejectionStatus1) SetAdditionalInformation(value string) {
	d.AdditionalInformation = (*Max350Text)(&value)
}
