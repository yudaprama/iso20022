package model

// Corporate action event cancellation status and reason.
type CorporateActionCancellation3 struct {

	// Specifies reasons for cancellation of a corporate action event.
	CancellationReasonCode *CorporateActionCancellationReason1Code `xml:"CxlRsnCd"`

	// Additional information about cancellation of a corporate action event.
	CancellationReason *Max140Text `xml:"CxlRsn,omitempty"`

	// Specifies the status of the details of the event.
	ProcessingStatus *CorporateActionEventStatus1 `xml:"PrcgSts"`
}

func (c *CorporateActionCancellation3) SetCancellationReasonCode(value string) {
	c.CancellationReasonCode = (*CorporateActionCancellationReason1Code)(&value)
}

func (c *CorporateActionCancellation3) SetCancellationReason(value string) {
	c.CancellationReason = (*Max140Text)(&value)
}

func (c *CorporateActionCancellation3) AddProcessingStatus() *CorporateActionEventStatus1 {
	c.ProcessingStatus = new(CorporateActionEventStatus1)
	return c.ProcessingStatus
}
