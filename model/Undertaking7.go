package model

// Information about an undertaking.
type Undertaking7 struct {

	// Unique and unambiguous identifier assigned by the issuer to the undertaking, for example the guarantee or standby number.
	Identification *Max35Text `xml:"Id"`

	// Party that issues the undertaking.
	Issuer *PartyIdentification43 `xml:"Issr"`
}

func (u *Undertaking7) SetIdentification(value string) {
	u.Identification = (*Max35Text)(&value)
}

func (u *Undertaking7) AddIssuer() *PartyIdentification43 {
	u.Issuer = new(PartyIdentification43)
	return u.Issuer
}
