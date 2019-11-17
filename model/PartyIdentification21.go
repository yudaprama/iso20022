package model

// Party involved in the settlement chain.
type PartyIdentification21 struct {

	// Party that legally owns the account.
	PartyIdentification *PartyIdentification2Choice `xml:"PtyId"`

	// Specifies the reference of the transaction at the party identified in the sequence.
	ProcessingReference *Max35Text `xml:"PrcgRef,omitempty"`

	// Date and optionally the time, at which this transaction was processed by the party identified in this sequence.
	ProcessingDate *DateAndDateTimeChoice `xml:"PrcgDt,omitempty"`

	// Contact person and contact information.
	ContactPerson *ContactIdentification2 `xml:"CtctPrsn,omitempty"`
}

func (p *PartyIdentification21) AddPartyIdentification() *PartyIdentification2Choice {
	p.PartyIdentification = new(PartyIdentification2Choice)
	return p.PartyIdentification
}

func (p *PartyIdentification21) SetProcessingReference(value string) {
	p.ProcessingReference = (*Max35Text)(&value)
}

func (p *PartyIdentification21) AddProcessingDate() *DateAndDateTimeChoice {
	p.ProcessingDate = new(DateAndDateTimeChoice)
	return p.ProcessingDate
}

func (p *PartyIdentification21) AddContactPerson() *ContactIdentification2 {
	p.ContactPerson = new(ContactIdentification2)
	return p.ContactPerson
}
