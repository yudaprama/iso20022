package model

// Provides reason of the rejection of an election amendment request.
type CorporateActionAmendmentRejectionStatus1 struct {

	// The rejection reason.
	Reason []*RejectionReason8FormatChoice `xml:"Rsn"`

	// Additional information about the status.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionAmendmentRejectionStatus1) AddReason() *RejectionReason8FormatChoice {
	newValue := new(RejectionReason8FormatChoice)
	c.Reason = append(c.Reason, newValue)
	return newValue
}

func (c *CorporateActionAmendmentRejectionStatus1) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}
