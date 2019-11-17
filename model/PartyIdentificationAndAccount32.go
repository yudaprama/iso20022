package model

// Party involved in the settlement chain.
type PartyIdentificationAndAccount32 struct {

	// Party that legally owns the account.
	PartyIdentification *PartyIdentification33Choice `xml:"PtyId"`

	// Identification of the account owned by the party.
	AccountIdentification *Max35Text `xml:"AcctId,omitempty"`

	// Specifies the reference of the transaction at the party identified in the sequence.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Date and optionally the time, at which this transaction was processed by the party identified in this sequence.
	ProcessingDate *DateAndDateTimeChoice `xml:"PrcgDt,omitempty"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	SubAccountDetails *SubAccount1 `xml:"SubAcctDtls,omitempty"`

	// Contact person and contact information.
	ContactPerson *ContactIdentification2 `xml:"CtctPrsn,omitempty"`
}

func (p *PartyIdentificationAndAccount32) AddPartyIdentification() *PartyIdentification33Choice {
	p.PartyIdentification = new(PartyIdentification33Choice)
	return p.PartyIdentification
}

func (p *PartyIdentificationAndAccount32) SetAccountIdentification(value string) {
	p.AccountIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount32) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount32) AddProcessingDate() *DateAndDateTimeChoice {
	p.ProcessingDate = new(DateAndDateTimeChoice)
	return p.ProcessingDate
}

func (p *PartyIdentificationAndAccount32) AddSubAccountDetails() *SubAccount1 {
	p.SubAccountDetails = new(SubAccount1)
	return p.SubAccountDetails
}

func (p *PartyIdentificationAndAccount32) AddContactPerson() *ContactIdentification2 {
	p.ContactPerson = new(ContactIdentification2)
	return p.ContactPerson
}
