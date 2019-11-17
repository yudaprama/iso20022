package model

// Reason for a conditionally accepted status.
type ConditionallyAcceptedStatusReason1 struct {

	// Reason for a conditionally accepted status in structured form.
	Structured []*ConditionallyAcceptedStatusReason1Code `xml:"Strd"`

	// Reason for a conditionally accepted status in free format text.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (c *ConditionallyAcceptedStatusReason1) AddStructured(value string) {
	c.Structured = append(c.Structured, (*ConditionallyAcceptedStatusReason1Code)(&value))
}

func (c *ConditionallyAcceptedStatusReason1) SetAdditionalInformation(value string) {
	c.AdditionalInformation = (*Max350Text)(&value)
}
