package model

// Entity involved in an activity.
type PartyIdentification1 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max70Text `xml:"Nm,omitempty"`

	// Postal address of a party.
	PostalAddress *PostalAddress1 `xml:"PstlAdr,omitempty"`

	// Unique and unambiguous identification of a party.
	Identification *Party1Choice `xml:"Id,omitempty"`
}

func (p *PartyIdentification1) SetName(value string) {
	p.Name = (*Max70Text)(&value)
}

func (p *PartyIdentification1) AddPostalAddress() *PostalAddress1 {
	p.PostalAddress = new(PostalAddress1)
	return p.PostalAddress
}

func (p *PartyIdentification1) AddIdentification() *Party1Choice {
	p.Identification = new(Party1Choice)
	return p.Identification
}
