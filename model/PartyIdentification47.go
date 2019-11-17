package model

// Provides information about identification of the party .
type PartyIdentification47 struct {

	// Identification of a party.
	Identification *PartyIdentification39Choice `xml:"Id"`

	// Reference meaningful to the party identified.
	ProcessingIdentification *Max35Text `xml:"PrcgId,omitempty"`

	// Provides additional information regarding the party, for example, the contact unit or person responsible for the transaction identified in the message.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`

	// Provides alternate identification for a party using an id type, a country code and a text field.
	AlternateIdentification []*AlternatePartyIdentification2 `xml:"AltrnId,omitempty"`
}

func (p *PartyIdentification47) AddIdentification() *PartyIdentification39Choice {
	p.Identification = new(PartyIdentification39Choice)
	return p.Identification
}

func (p *PartyIdentification47) SetProcessingIdentification(value string) {
	p.ProcessingIdentification = (*Max35Text)(&value)
}

func (p *PartyIdentification47) SetAdditionalInformation(value string) {
	p.AdditionalInformation = (*Max350Text)(&value)
}

func (p *PartyIdentification47) AddAlternateIdentification() *AlternatePartyIdentification2 {
	newValue := new(AlternatePartyIdentification2)
	p.AlternateIdentification = append(p.AlternateIdentification, newValue)
	return newValue
}
