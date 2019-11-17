package model

// Party involved in the settlement chain.
type PartyIdentification97 struct {

	// Party that legally owns the account.
	PartyIdentification *PartyIdentification70Choice `xml:"PtyId"`

	// Specifies the reference of the transaction at the party identified in the sequence.
	ProcessingReference *Max35Text `xml:"PrcgRef,omitempty"`

	// Date and optionally the time, at which this transaction was processed by the party identified in this sequence.
	ProcessingDate *DateAndDateTimeChoice `xml:"PrcgDt,omitempty"`

	// Contact person and contact information.
	ContactPerson *ContactIdentification2 `xml:"CtctPrsn,omitempty"`
}

func (p *PartyIdentification97) AddPartyIdentification() *PartyIdentification70Choice {
	p.PartyIdentification = new(PartyIdentification70Choice)
	return p.PartyIdentification
}

func (p *PartyIdentification97) SetProcessingReference(value string) {
	p.ProcessingReference = (*Max35Text)(&value)
}

func (p *PartyIdentification97) AddProcessingDate() *DateAndDateTimeChoice {
	p.ProcessingDate = new(DateAndDateTimeChoice)
	return p.ProcessingDate
}

func (p *PartyIdentification97) AddContactPerson() *ContactIdentification2 {
	p.ContactPerson = new(ContactIdentification2)
	return p.ContactPerson
}
