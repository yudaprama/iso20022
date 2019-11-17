package model

// Provides reason of the rejection of a movement.
type CorporateActionMovementRejectionStatus1 struct {

	// The rejection reason.
	Reason []*RejectionReason13FormatChoice `xml:"Rsn"`

	// Additional information about the status.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionMovementRejectionStatus1) AddReason() *RejectionReason13FormatChoice {
	newValue := new(RejectionReason13FormatChoice)
	c.Reason = append(c.Reason, newValue)
	return newValue
}

func (c *CorporateActionMovementRejectionStatus1) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}
