package model

// Party involved in the settlement chain.
type PartyIdentificationAndAccount102 struct {

	// Party that legally owns the account.
	PartyIdentification *PartyIdentification33Choice `xml:"PtyId"`

	// Identification of the account owned by the party.
	AccountIdentification *Max35Text `xml:"AcctId,omitempty"`

	// Specifies the reference of the transaction at the party identified in the sequence.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Date and optionally the time, at which this transaction was processed by the party identified in this sequence.
	ProcessingDate *DateAndDateTimeChoice `xml:"PrcgDt,omitempty"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	SubAccount *SubAccount4 `xml:"SubAcct,omitempty"`

	// Contact person and contact information.
	ContactPerson *ContactIdentification2 `xml:"CtctPrsn,omitempty"`
}

func (p *PartyIdentificationAndAccount102) AddPartyIdentification() *PartyIdentification33Choice {
	p.PartyIdentification = new(PartyIdentification33Choice)
	return p.PartyIdentification
}

func (p *PartyIdentificationAndAccount102) SetAccountIdentification(value string) {
	p.AccountIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount102) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount102) AddProcessingDate() *DateAndDateTimeChoice {
	p.ProcessingDate = new(DateAndDateTimeChoice)
	return p.ProcessingDate
}

func (p *PartyIdentificationAndAccount102) AddSubAccount() *SubAccount4 {
	p.SubAccount = new(SubAccount4)
	return p.SubAccount
}

func (p *PartyIdentificationAndAccount102) AddContactPerson() *ContactIdentification2 {
	p.ContactPerson = new(ContactIdentification2)
	return p.ContactPerson
}
