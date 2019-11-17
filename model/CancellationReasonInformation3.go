package model

// Set of elements used to provide information on the reason of the cancellation request.
type CancellationReasonInformation3 struct {

	// Party that issues the cancellation request.
	Originator *PartyIdentification32 `xml:"Orgtr,omitempty"`

	// Specifies the reason for the cancellation.
	Reason *CancellationReason2Choice `xml:"Rsn,omitempty"`

	// Further details on the cancellation request reason.
	AdditionalInformation []*Max105Text `xml:"AddtlInf,omitempty"`
}

func (c *CancellationReasonInformation3) AddOriginator() *PartyIdentification32 {
	c.Originator = new(PartyIdentification32)
	return c.Originator
}

func (c *CancellationReasonInformation3) AddReason() *CancellationReason2Choice {
	c.Reason = new(CancellationReason2Choice)
	return c.Reason
}

func (c *CancellationReasonInformation3) AddAdditionalInformation(value string) {
	c.AdditionalInformation = append(c.AdditionalInformation, (*Max105Text)(&value))
}
