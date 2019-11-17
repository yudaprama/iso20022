package model

// Reason for a pending status.
type PendingStatusReason14 struct {

	// Reason for the pending account status.
	Code *PendingStatusReason2Choice `xml:"Cd"`

	// Additional information about the reason for the pending account status.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (p *PendingStatusReason14) AddCode() *PendingStatusReason2Choice {
	p.Code = new(PendingStatusReason2Choice)
	return p.Code
}

func (p *PendingStatusReason14) SetAdditionalInformation(value string) {
	p.AdditionalInformation = (*Max350Text)(&value)
}
