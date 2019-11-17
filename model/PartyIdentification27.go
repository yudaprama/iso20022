package model

// Entity involved in an activity.
type PartyIdentification27 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max70Text `xml:"Nm"`

	// Unique and unambiguous identifier assigned to a party using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification4 `xml:"PrtryId,omitempty"`

	// Specifies the country of the party.
	Country *CountryCode `xml:"Ctry"`
}

func (p *PartyIdentification27) SetName(value string) {
	p.Name = (*Max70Text)(&value)
}

func (p *PartyIdentification27) AddProprietaryIdentification() *GenericIdentification4 {
	p.ProprietaryIdentification = new(GenericIdentification4)
	return p.ProprietaryIdentification
}

func (p *PartyIdentification27) SetCountry(value string) {
	p.Country = (*CountryCode)(&value)
}
