package model

// Party involved in the settlement chain.
type PartyIdentificationAndAccount124 struct {

	// Party that legally owns the account.
	PartyIdentification *PartyIdentification70Choice `xml:"PtyId"`

	// Identification of the account owned by the party.
	AccountIdentification *Max35Text `xml:"AcctId,omitempty"`

	// Specifies the reference of the transaction at the party identified in the sequence.
	ProcessingReference *Max35Text `xml:"PrcgRef,omitempty"`

	// Date and optionally the time, at which this transaction was processed by the party identified in this sequence.
	ProcessingDate *DateAndDateTimeChoice `xml:"PrcgDt,omitempty"`
}

func (p *PartyIdentificationAndAccount124) AddPartyIdentification() *PartyIdentification70Choice {
	p.PartyIdentification = new(PartyIdentification70Choice)
	return p.PartyIdentification
}

func (p *PartyIdentificationAndAccount124) SetAccountIdentification(value string) {
	p.AccountIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount124) SetProcessingReference(value string) {
	p.ProcessingReference = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount124) AddProcessingDate() *DateAndDateTimeChoice {
	p.ProcessingDate = new(DateAndDateTimeChoice)
	return p.ProcessingDate
}
