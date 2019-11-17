package model

// Party and account details.
type PartyIdentificationAndAccount45 struct {

	// Identification of the party.
	Identification *PartyIdentification45Choice `xml:"Id"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount13 `xml:"SfkpgAcct,omitempty"`

	// Date/time at which the instruction was processed by the specified party.
	ProcessingDate *DateAndDateTimeChoice `xml:"PrcgDt,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation1 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount45) AddIdentification() *PartyIdentification45Choice {
	p.Identification = new(PartyIdentification45Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount45) AddAlternateIdentification() *AlternatePartyIdentification2 {
	p.AlternateIdentification = new(AlternatePartyIdentification2)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount45) AddSafekeepingAccount() *SecuritiesAccount13 {
	p.SafekeepingAccount = new(SecuritiesAccount13)
	return p.SafekeepingAccount
}

func (p *PartyIdentificationAndAccount45) AddProcessingDate() *DateAndDateTimeChoice {
	p.ProcessingDate = new(DateAndDateTimeChoice)
	return p.ProcessingDate
}

func (p *PartyIdentificationAndAccount45) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentificationAndAccount45) AddAdditionalInformation() *PartyTextInformation1 {
	p.AdditionalInformation = new(PartyTextInformation1)
	return p.AdditionalInformation
}
