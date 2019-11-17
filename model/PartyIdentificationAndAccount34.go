package model

// Entity involved in an activity.
type PartyIdentificationAndAccount34 struct {

	// Identification of a party.
	Identification *PartyIdentification32Choice `xml:"Id"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification5 `xml:"AltrnId,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct,omitempty"`
}

func (p *PartyIdentificationAndAccount34) AddIdentification() *PartyIdentification32Choice {
	p.Identification = new(PartyIdentification32Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount34) SetAdditionalInformation(value string) {
	p.AdditionalInformation = (*Max350Text)(&value)
}

func (p *PartyIdentificationAndAccount34) AddAlternateIdentification() *AlternatePartyIdentification5 {
	p.AlternateIdentification = new(AlternatePartyIdentification5)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount34) SetSafekeepingAccount(value string) {
	p.SafekeepingAccount = (*Max35Text)(&value)
}
