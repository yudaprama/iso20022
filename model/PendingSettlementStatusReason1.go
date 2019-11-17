package model

// Reason for a  pending settlement status.
type PendingSettlementStatusReason1 struct {

	// Reason for a settlement pending status in structured form.
	Structured *PendingSettlementStatusReason1Code `xml:"Strd"`

	// Additional information about the reason for the settlement pending status in textual form.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (p *PendingSettlementStatusReason1) SetStructured(value string) {
	p.Structured = (*PendingSettlementStatusReason1Code)(&value)
}

func (p *PendingSettlementStatusReason1) SetAdditionalInformation(value string) {
	p.AdditionalInformation = (*Max350Text)(&value)
}
