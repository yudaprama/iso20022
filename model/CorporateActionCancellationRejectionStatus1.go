package model

// Provides reason of the rejection of an election cancellation request.
type CorporateActionCancellationRejectionStatus1 struct {

	// The rejection reason.
	Reason []*RejectionReason9FormatChoice `xml:"Rsn"`

	// Additional information about the status.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionCancellationRejectionStatus1) AddReason() *RejectionReason9FormatChoice {
	newValue := new(RejectionReason9FormatChoice)
	c.Reason = append(c.Reason, newValue)
	return newValue
}

func (c *CorporateActionCancellationRejectionStatus1) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}
