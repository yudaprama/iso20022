package model

// Party and related security certificate.
type PartyAndCertificate3 struct {

	// Specifies the type of change.
	ModificationCode *Modification1Code `xml:"ModCd,omitempty"`

	// Entity involved in an activity.
	Party *PartyIdentification43 `xml:"Pty"`

	// Security certificate used to sign electronically.
	Certificate *Max10KBinary `xml:"Cert,omitempty"`
}

func (p *PartyAndCertificate3) SetModificationCode(value string) {
	p.ModificationCode = (*Modification1Code)(&value)
}

func (p *PartyAndCertificate3) AddParty() *PartyIdentification43 {
	p.Party = new(PartyIdentification43)
	return p.Party
}

func (p *PartyAndCertificate3) SetCertificate(value string) {
	p.Certificate = (*Max10KBinary)(&value)
}
