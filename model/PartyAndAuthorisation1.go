package model

// Party and related authorisation.
type PartyAndAuthorisation1 struct {

	// Specifies a party or a group of parties.
	PartyOrGroup *PartyOrGroup1Choice `xml:"PtyOrGrp"`

	// Order in which the mandate holder has to sign.
	SignatureOrder *Max15PlusSignedNumericText `xml:"SgntrOrdr,omitempty"`

	// Authorisation granted to a mandate holder.
	Authorisation *Authorisation2 `xml:"Authstn"`
}

func (p *PartyAndAuthorisation1) AddPartyOrGroup() *PartyOrGroup1Choice {
	p.PartyOrGroup = new(PartyOrGroup1Choice)
	return p.PartyOrGroup
}

func (p *PartyAndAuthorisation1) SetSignatureOrder(value string) {
	p.SignatureOrder = (*Max15PlusSignedNumericText)(&value)
}

func (p *PartyAndAuthorisation1) AddAuthorisation() *Authorisation2 {
	p.Authorisation = new(Authorisation2)
	return p.Authorisation
}
