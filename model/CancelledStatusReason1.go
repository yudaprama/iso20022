package model

// Reason for a cancelled status.
type CancelledStatusReason1 struct {

	// Reason for a cancelled status in free format text.
	AdditionalInformation *Max350Text `xml:"AddtlInf"`
}

func (c *CancelledStatusReason1) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}
