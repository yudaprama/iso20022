package model

// Status is cancelled.
type CancellationStatus22Choice struct {

	// Status of the order cancellation request is cancelled.
	Status *OrderCancellationStatus2Code `xml:"Sts"`

	// Status of the order cancellation request is rejected.
	Rejected *RejectedStatus10 `xml:"Rjctd"`
}

func (c *CancellationStatus22Choice) SetStatus(value string) {
	c.Status = (*OrderCancellationStatus2Code)(&value)
}

func (c *CancellationStatus22Choice) AddRejected() *RejectedStatus10 {
	c.Rejected = new(RejectedStatus10)
	return c.Rejected
}
