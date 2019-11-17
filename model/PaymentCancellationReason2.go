package model

// Provides further details on the reason of the cancellation request.
type PaymentCancellationReason2 struct {

	// Party that issues the cancellation request.
	Originator *PartyIdentification43 `xml:"Orgtr,omitempty"`

	// Specifies the reason for the cancellation.
	Reason *CancellationReason14Choice `xml:"Rsn,omitempty"`

	// Further details on the cancellation request reason.
	AdditionalInformation []*Max105Text `xml:"AddtlInf,omitempty"`
}

func (p *PaymentCancellationReason2) AddOriginator() *PartyIdentification43 {
	p.Originator = new(PartyIdentification43)
	return p.Originator
}

func (p *PaymentCancellationReason2) AddReason() *CancellationReason14Choice {
	p.Reason = new(CancellationReason14Choice)
	return p.Reason
}

func (p *PaymentCancellationReason2) AddAdditionalInformation(value string) {
	p.AdditionalInformation = append(p.AdditionalInformation, (*Max105Text)(&value))
}
