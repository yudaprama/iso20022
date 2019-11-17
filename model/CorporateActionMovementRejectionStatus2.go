package model

// Provides information about the rejection of a movement cancellation request.
type CorporateActionMovementRejectionStatus2 struct {

	// The rejection reason.
	Reason []*RejectionReason14FormatChoice `xml:"Rsn"`

	// Additional information about the status.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionMovementRejectionStatus2) AddReason() *RejectionReason14FormatChoice {
	newValue := new(RejectionReason14FormatChoice)
	c.Reason = append(c.Reason, newValue)
	return newValue
}

func (c *CorporateActionMovementRejectionStatus2) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}
