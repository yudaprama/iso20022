package model

// Set of elements used to provide information on the status of the cancellation request.
type CancellationStatusReasonInformation1 struct {

	// Party that issues the cancellation status.
	Originator *PartyIdentification32 `xml:"Orgtr,omitempty"`

	// Specifies the reason for the status report.
	Reason *CancellationStatusReason1Choice `xml:"Rsn,omitempty"`

	// Further details on the cancellation status reason.
	AdditionalInformation []*Max105Text `xml:"AddtlInf,omitempty"`
}

func (c *CancellationStatusReasonInformation1) AddOriginator() *PartyIdentification32 {
	c.Originator = new(PartyIdentification32)
	return c.Originator
}

func (c *CancellationStatusReasonInformation1) AddReason() *CancellationStatusReason1Choice {
	c.Reason = new(CancellationStatusReason1Choice)
	return c.Reason
}

func (c *CancellationStatusReasonInformation1) AddAdditionalInformation(value string) {
	c.AdditionalInformation = append(c.AdditionalInformation, (*Max105Text)(&value))
}
