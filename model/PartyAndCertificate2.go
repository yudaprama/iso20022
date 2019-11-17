package model

// Party and related security certificate.
type PartyAndCertificate2 struct {

	// Entity involved in an activity.
	Party *PartyIdentification43 `xml:"Pty"`

	// Security certificate used to sign electronically.
	Certificate *Max10KBinary `xml:"Cert,omitempty"`
}

func (p *PartyAndCertificate2) AddParty() *PartyIdentification43 {
	p.Party = new(PartyIdentification43)
	return p.Party
}

func (p *PartyAndCertificate2) SetCertificate(value string) {
	p.Certificate = (*Max10KBinary)(&value)
}
