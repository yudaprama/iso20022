package model

// Further information on the cancellation reason of the transaction.
type CancellationReasonInformation1 struct {

	// Party issuing the cancellation request.
	CancellationOriginator *PartyIdentification8 `xml:"CxlOrgtr,omitempty"`

	// Specifies the reason for the cancellation.
	CancellationReason *CancellationReason1Choice `xml:"CxlRsn,omitempty"`

	// Further details on the cancellation request reason.
	//
	// Usage: Additional cancellation reason information can be used to further detail the cancellation request reason.
	AdditionalCancellationReasonInformation []*Max105Text `xml:"AddtlCxlRsnInf,omitempty"`
}

func (c *CancellationReasonInformation1) AddCancellationOriginator() *PartyIdentification8 {
	c.CancellationOriginator = new(PartyIdentification8)
	return c.CancellationOriginator
}

func (c *CancellationReasonInformation1) AddCancellationReason() *CancellationReason1Choice {
	c.CancellationReason = new(CancellationReason1Choice)
	return c.CancellationReason
}

func (c *CancellationReasonInformation1) AddAdditionalCancellationReasonInformation(value string) {
	c.AdditionalCancellationReasonInformation = append(c.AdditionalCancellationReasonInformation, (*Max105Text)(&value))
}
