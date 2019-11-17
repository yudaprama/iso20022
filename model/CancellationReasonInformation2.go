package model

// Set of elements used to provide information on the reason of the mandate cancellation request.
type CancellationReasonInformation2 struct {

	// Party that issues the cancellation request.
	Originator *PartyIdentification32 `xml:"Orgtr,omitempty"`

	// Specifies the reason for the cancellation request.
	Reason *MandateReason1Choice `xml:"Rsn"`

	// Further details on the cancellation request reason.
	AdditionalInformation []*Max105Text `xml:"AddtlInf,omitempty"`
}

func (c *CancellationReasonInformation2) AddOriginator() *PartyIdentification32 {
	c.Originator = new(PartyIdentification32)
	return c.Originator
}

func (c *CancellationReasonInformation2) AddReason() *MandateReason1Choice {
	c.Reason = new(MandateReason1Choice)
	return c.Reason
}

func (c *CancellationReasonInformation2) AddAdditionalInformation(value string) {
	c.AdditionalInformation = append(c.AdditionalInformation, (*Max105Text)(&value))
}
