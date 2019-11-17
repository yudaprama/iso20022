package model

// Reason for a pending account opening status.
type PendingOpeningStatusReason1 struct {

	// Reason for the pending account opening status.
	Code *PendingOpeningStatusReason2Choice `xml:"Cd"`

	// Additional information about the reason for the pending account opening status.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (p *PendingOpeningStatusReason1) AddCode() *PendingOpeningStatusReason2Choice {
	p.Code = new(PendingOpeningStatusReason2Choice)
	return p.Code
}

func (p *PendingOpeningStatusReason1) SetAdditionalInformation(value string) {
	p.AdditionalInformation = (*Max350Text)(&value)
}
