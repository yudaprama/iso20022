package model

// Provides further details on the status of the cancellation request.
type CancellationStatusReason2 struct {

	// Party that issues the cancellation status.
	Originator *PartyIdentification43 `xml:"Orgtr,omitempty"`

	// Specifies the reason for the status report.
	Reason *CancellationStatusReason2Choice `xml:"Rsn,omitempty"`

	// Further details on the cancellation status reason.
	AdditionalInformation []*Max105Text `xml:"AddtlInf,omitempty"`
}

func (c *CancellationStatusReason2) AddOriginator() *PartyIdentification43 {
	c.Originator = new(PartyIdentification43)
	return c.Originator
}

func (c *CancellationStatusReason2) AddReason() *CancellationStatusReason2Choice {
	c.Reason = new(CancellationStatusReason2Choice)
	return c.Reason
}

func (c *CancellationStatusReason2) AddAdditionalInformation(value string) {
	c.AdditionalInformation = append(c.AdditionalInformation, (*Max105Text)(&value))
}
