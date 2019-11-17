package model

// Specification of the cancellation right.
type CancellationRight1 struct {

	// Structured format.
	Structured *CancellationRight2Code `xml:"Strd"`

	// Additional information about the type of charge.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *CancellationRight1) SetStructured(value string) {
	c.Structured = (*CancellationRight2Code)(&value)
}

func (c *CancellationRight1) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}
