package model

// Reason for a partially settled status.
type PartiallySettledStatus10 struct {

	// Reason for the partially settled status.
	Reason *PartiallySettled21Choice `xml:"Rsn"`

	// Additional information about the partially settled reason.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (p *PartiallySettledStatus10) AddReason() *PartiallySettled21Choice {
	p.Reason = new(PartiallySettled21Choice)
	return p.Reason
}

func (p *PartiallySettledStatus10) SetAdditionalInformation(value string) {
	p.AdditionalInformation = (*Max350Text)(&value)
}
