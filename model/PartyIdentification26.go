package model

// Entity involved in an activity.
type PartyIdentification26 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max70Text `xml:"Nm"`

	// Unique and unambiguous identifier assigned to a party using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification4 `xml:"PrtryId,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress *PostalAddress5 `xml:"PstlAdr"`
}

func (p *PartyIdentification26) SetName(value string) {
	p.Name = (*Max70Text)(&value)
}

func (p *PartyIdentification26) AddProprietaryIdentification() *GenericIdentification4 {
	p.ProprietaryIdentification = new(GenericIdentification4)
	return p.ProprietaryIdentification
}

func (p *PartyIdentification26) AddPostalAddress() *PostalAddress5 {
	p.PostalAddress = new(PostalAddress5)
	return p.PostalAddress
}
