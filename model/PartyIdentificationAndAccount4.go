package model

// Party involved in the settlement chain.
type PartyIdentificationAndAccount4 struct {

	// Party that legally owns the account.
	PartyIdentification *PartyIdentification2Choice `xml:"PtyId"`

	// Identification of the account owned by the party.
	AccountIdentification *AccountIdentification1 `xml:"AcctId,omitempty"`

	// Specifies the reference of the transaction at the party identified in the sequence.
	ProcessingReference *Max35Text `xml:"PrcgRef,omitempty"`

	// Date and optionally the time, at which this transaction was processed by the party identified in this sequence.
	ProcessingDate *DateAndDateTimeChoice `xml:"PrcgDt,omitempty"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	SubAccountDetails *SubAccount1 `xml:"SubAcctDtls,omitempty"`

	// Contact person and contact information.
	ContactPerson *ContactIdentification2 `xml:"CtctPrsn,omitempty"`
}

func (p *PartyIdentificationAndAccount4) AddPartyIdentification() *PartyIdentification2Choice {
	p.PartyIdentification = new(PartyIdentification2Choice)
	return p.PartyIdentification
}

func (p *PartyIdentificationAndAccount4) AddAccountIdentification() *AccountIdentification1 {
	p.AccountIdentification = new(AccountIdentification1)
	return p.AccountIdentification
}

func (p *PartyIdentificationAndAccount4) SetProcessingReference(value string) {
	p.ProcessingReference = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount4) AddProcessingDate() *DateAndDateTimeChoice {
	p.ProcessingDate = new(DateAndDateTimeChoice)
	return p.ProcessingDate
}

func (p *PartyIdentificationAndAccount4) AddSubAccountDetails() *SubAccount1 {
	p.SubAccountDetails = new(SubAccount1)
	return p.SubAccountDetails
}

func (p *PartyIdentificationAndAccount4) AddContactPerson() *ContactIdentification2 {
	p.ContactPerson = new(ContactIdentification2)
	return p.ContactPerson
}
