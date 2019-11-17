package model

// Entity involved in an activity.
type PartyAndCertificate1 struct {

	// Entity involved in an activity.
	Party *PartyIdentification41 `xml:"Pty"`

	// Security certificate used to sign electronically.
	Certificate *Max10KBinary `xml:"Cert,omitempty"`

	// Order in which the mandate holder has to sign.
	SignatureOrder *Max15PlusSignedNumericText `xml:"SgntrOrdr,omitempty"`

	// Authorisation granted to a mandate holder.
	Authorisation *Authorisation1 `xml:"Authstn"`
}

func (p *PartyAndCertificate1) AddParty() *PartyIdentification41 {
	p.Party = new(PartyIdentification41)
	return p.Party
}

func (p *PartyAndCertificate1) SetCertificate(value string) {
	p.Certificate = (*Max10KBinary)(&value)
}

func (p *PartyAndCertificate1) SetSignatureOrder(value string) {
	p.SignatureOrder = (*Max15PlusSignedNumericText)(&value)
}

func (p *PartyAndCertificate1) AddAuthorisation() *Authorisation1 {
	p.Authorisation = new(Authorisation1)
	return p.Authorisation
}
