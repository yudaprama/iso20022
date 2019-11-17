package model

// Provides reason of the rejection of a deactivation cancellation request.
type CorporateActionDeactivationCancellationRejectionStatus1 struct {

	// The rejection reason.
	Reason []*RejectionReason7FormatChoice `xml:"Rsn"`

	// Additional information about the status.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionDeactivationCancellationRejectionStatus1) AddReason() *RejectionReason7FormatChoice {
	newValue := new(RejectionReason7FormatChoice)
	c.Reason = append(c.Reason, newValue)
	return newValue
}

func (c *CorporateActionDeactivationCancellationRejectionStatus1) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}
