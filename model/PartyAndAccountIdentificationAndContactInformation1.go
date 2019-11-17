package model

// Entity involved in an activity.
type PartyAndAccountIdentificationAndContactInformation1 struct {

	// Identification of the party that legally owns the account.
	PartyIdentification *PartyIdentification8 `xml:"PtyId"`

	// Identification of the account owned by the party.
	AccountIdentification *CashAccount7 `xml:"AcctId,omitempty"`

	// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
	ContactInformation *ContactIdentification1 `xml:"CtctInf,omitempty"`
}

func (p *PartyAndAccountIdentificationAndContactInformation1) AddPartyIdentification() *PartyIdentification8 {
	p.PartyIdentification = new(PartyIdentification8)
	return p.PartyIdentification
}

func (p *PartyAndAccountIdentificationAndContactInformation1) AddAccountIdentification() *CashAccount7 {
	p.AccountIdentification = new(CashAccount7)
	return p.AccountIdentification
}

func (p *PartyAndAccountIdentificationAndContactInformation1) AddContactInformation() *ContactIdentification1 {
	p.ContactInformation = new(ContactIdentification1)
	return p.ContactInformation
}
