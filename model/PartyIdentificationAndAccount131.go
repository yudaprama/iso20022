package model

// Party and account details.
type PartyIdentificationAndAccount131 struct {

	// Identification of the party.
	Identification *PartyIdentification104Choice `xml:"Id"`

	// Legal Entity Identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`

	// Alternate identification for a party.
	AlternateIdentification *AlternatePartyIdentification9 `xml:"AltrnId,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount27 `xml:"SfkpgAcct,omitempty"`

	// Date/time at which the instruction was processed by the specified party.
	ProcessingDate *DateAndDateTimeChoice `xml:"PrcgDt,omitempty"`

	// Unambiguous identification of the transaction for the party identified.
	ProcessingIdentification *RestrictedFINXMax16Text `xml:"PrcgId,omitempty"`

	// Provides additional information to a party identification.
	AdditionalInformation *PartyTextInformation3 `xml:"AddtlInf,omitempty"`
}

func (p *PartyIdentificationAndAccount131) AddIdentification() *PartyIdentification104Choice {
	p.Identification = new(PartyIdentification104Choice)
	return p.Identification
}

func (p *PartyIdentificationAndAccount131) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}

func (p *PartyIdentificationAndAccount131) AddAlternateIdentification() *AlternatePartyIdentification9 {
	p.AlternateIdentification = new(AlternatePartyIdentification9)
	return p.AlternateIdentification
}

func (p *PartyIdentificationAndAccount131) AddSafekeepingAccount() *SecuritiesAccount27 {
	p.SafekeepingAccount = new(SecuritiesAccount27)
	return p.SafekeepingAccount
}

func (p *PartyIdentificationAndAccount131) AddProcessingDate() *DateAndDateTimeChoice {
	p.ProcessingDate = new(DateAndDateTimeChoice)
	return p.ProcessingDate
}

func (p *PartyIdentificationAndAccount131) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (p *PartyIdentificationAndAccount131) AddAdditionalInformation() *PartyTextInformation3 {
	p.AdditionalInformation = new(PartyTextInformation3)
	return p.AdditionalInformation
}
