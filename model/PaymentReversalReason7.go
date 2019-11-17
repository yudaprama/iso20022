package model

// Provides further details on the reason of the reversal of the transaction.
type PaymentReversalReason7 struct {

	// Party that issues the reversal.
	Originator *PartyIdentification43 `xml:"Orgtr,omitempty"`

	// Specifies the reason for the reversal.
	Reason *ReversalReason4Choice `xml:"Rsn,omitempty"`

	// Further details on the reversal reason.
	AdditionalInformation []*Max105Text `xml:"AddtlInf,omitempty"`
}

func (p *PaymentReversalReason7) AddOriginator() *PartyIdentification43 {
	p.Originator = new(PartyIdentification43)
	return p.Originator
}

func (p *PaymentReversalReason7) AddReason() *ReversalReason4Choice {
	p.Reason = new(ReversalReason4Choice)
	return p.Reason
}

func (p *PaymentReversalReason7) AddAdditionalInformation(value string) {
	p.AdditionalInformation = append(p.AdditionalInformation, (*Max105Text)(&value))
}
