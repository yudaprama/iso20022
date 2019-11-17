package model

// Provides reason of the rejection of an information advice.
type CorporateActionInformationRejectedStatus1 struct {

	// The rejection reason.
	Reason []*RejectionReason15FormatChoice `xml:"Rsn"`

	// Additional information about the status.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionInformationRejectedStatus1) AddReason() *RejectionReason15FormatChoice {
	newValue := new(RejectionReason15FormatChoice)
	c.Reason = append(c.Reason, newValue)
	return newValue
}

func (c *CorporateActionInformationRejectedStatus1) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}
