package model

// Choice of status for the cancellation processing.
type CancellationProcessingStatus6Choice struct {

	// Trade is in cancelation pending.
	CancellationPending *CancellationReason11Choice `xml:"CxlPdg"`

	// Cancelation request for the trade..
	CancellationRequested *ProprietaryReason1 `xml:"CxlReqd"`

	// Cancellation is completed.
	CancellationCompleted *ProprietaryReason1 `xml:"CxlCmpltd"`

	// Provides a proprietary status and a proprietary reason of the processing status of the trade.
	ProprietaryStatus *ProprietaryStatusAndReason1 `xml:"PrtrySts,omitempty"`
}

func (c *CancellationProcessingStatus6Choice) AddCancellationPending() *CancellationReason11Choice {
	c.CancellationPending = new(CancellationReason11Choice)
	return c.CancellationPending
}

func (c *CancellationProcessingStatus6Choice) AddCancellationRequested() *ProprietaryReason1 {
	c.CancellationRequested = new(ProprietaryReason1)
	return c.CancellationRequested
}

func (c *CancellationProcessingStatus6Choice) AddCancellationCompleted() *ProprietaryReason1 {
	c.CancellationCompleted = new(ProprietaryReason1)
	return c.CancellationCompleted
}

func (c *CancellationProcessingStatus6Choice) AddProprietaryStatus() *ProprietaryStatusAndReason1 {
	c.ProprietaryStatus = new(ProprietaryStatusAndReason1)
	return c.ProprietaryStatus
}
