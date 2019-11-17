package model

// Reason for a cancelled completed status.
type CancelledCompleteReason1 struct {

	// Reason for the cancelled complete status.
	Reason *CancellationCompleteReason1Choice `xml:"Rsn"`

	// Additional information about the cancelled complete status reason.
	AdditionalReasonInformation *Max350Text `xml:"AddtlRsnInf,omitempty"`
}

func (c *CancelledCompleteReason1) AddReason() *CancellationCompleteReason1Choice {
	c.Reason = new(CancellationCompleteReason1Choice)
	return c.Reason
}

func (c *CancelledCompleteReason1) SetAdditionalReasonInformation(value string) {
	c.AdditionalReasonInformation = (*Max350Text)(&value)
}
