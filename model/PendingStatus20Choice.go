package model

// Choice of format for the matching status.
type PendingStatus20Choice struct {

	// Allegement is forwarded.
	Forwarded *NoSpecifiedReason1 `xml:"Fwdd"`

	// Allegement is under investigation.
	UnderInvestigation *NoSpecifiedReason1 `xml:"UdrInvstgtn"`
}

func (p *PendingStatus20Choice) AddForwarded() *NoSpecifiedReason1 {
	p.Forwarded = new(NoSpecifiedReason1)
	return p.Forwarded
}

func (p *PendingStatus20Choice) AddUnderInvestigation() *NoSpecifiedReason1 {
	p.UnderInvestigation = new(NoSpecifiedReason1)
	return p.UnderInvestigation
}
