package model

// Corporate action event cancellation status and reason.
type CorporateActionCancellation1 struct {

	// Specifies reasons for cancellation of a corporate action event.
	CancellationReasonCode *CorporateActionCancellationReason1Code `xml:"CxlRsnCd"`

	// Additional information about cancellation of a corporate action event.
	CancellationReason *Max140Text `xml:"CxlRsn,omitempty"`

	// Specifies the status of the details of the event.
	ProcessingStatus *CorporateActionProcessingStatus1Choice `xml:"PrcgSts"`
}

func (c *CorporateActionCancellation1) SetCancellationReasonCode(value string) {
	c.CancellationReasonCode = (*CorporateActionCancellationReason1Code)(&value)
}

func (c *CorporateActionCancellation1) SetCancellationReason(value string) {
	c.CancellationReason = (*Max140Text)(&value)
}

func (c *CorporateActionCancellation1) AddProcessingStatus() *CorporateActionProcessingStatus1Choice {
	c.ProcessingStatus = new(CorporateActionProcessingStatus1Choice)
	return c.ProcessingStatus
}
