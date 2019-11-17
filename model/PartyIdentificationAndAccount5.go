package model

// Party involved in the settlement chain.
type PartyIdentificationAndAccount5 struct {

	// Party that legally owns the account.
	PartyIdentification *PartyIdentification2Choice `xml:"PtyId"`

	// Identification of the account owned by the party.
	AccountIdentification *AccountIdentification1 `xml:"AcctId,omitempty"`

	// Specifies the reference of the transaction at the party identified in the sequence.
	ProcessingReference *Max35Text `xml:"PrcgRef,omitempty"`

	// Date and optionally the time, at which this transaction was processed by the party identified in this sequence.
	ProcessingDate *DateAndDateTimeChoice `xml:"PrcgDt,omitempty"`
}

func (p *PartyIdentificationAndAccount5) AddPartyIdentification() *PartyIdentification2Choice {
	p.PartyIdentification = new(PartyIdentification2Choice)
	return p.PartyIdentification
}

func (p *PartyIdentificationAndAccount5) AddAccountIdentification() *AccountIdentification1 {
	p.AccountIdentification = new(AccountIdentification1)
	return p.AccountIdentification
}

func (p *PartyIdentificationAndAccount5) SetProcessingReference(value string) {
	p.ProcessingReference = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount5) AddProcessingDate() *DateAndDateTimeChoice {
	p.ProcessingDate = new(DateAndDateTimeChoice)
	return p.ProcessingDate
}
