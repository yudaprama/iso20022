package model

// Party and related authorisation.
type PartyAndAuthorisation3 struct {

	// Specifies the type of change.
	ModificationCode *Modification1Code `xml:"ModCd,omitempty"`

	// Specifies a party or a group of parties.
	PartyOrGroup *PartyOrGroup1Choice `xml:"PtyOrGrp"`

	// Order in which the mandate holder has to sign.
	SignatureOrder *Max15PlusSignedNumericText `xml:"SgntrOrdr,omitempty"`

	// Authorisation granted to a mandate holder.
	Authorisation *Authorisation2 `xml:"Authstn"`
}

func (p *PartyAndAuthorisation3) SetModificationCode(value string) {
	p.ModificationCode = (*Modification1Code)(&value)
}

func (p *PartyAndAuthorisation3) AddPartyOrGroup() *PartyOrGroup1Choice {
	p.PartyOrGroup = new(PartyOrGroup1Choice)
	return p.PartyOrGroup
}

func (p *PartyAndAuthorisation3) SetSignatureOrder(value string) {
	p.SignatureOrder = (*Max15PlusSignedNumericText)(&value)
}

func (p *PartyAndAuthorisation3) AddAuthorisation() *Authorisation2 {
	p.Authorisation = new(Authorisation2)
	return p.Authorisation
}
