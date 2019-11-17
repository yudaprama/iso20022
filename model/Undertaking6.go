package model

// Information about an undertaking.
type Undertaking6 struct {

	// Unique and unambiguous identifier assigned by the issuer to the undertaking, for example the guarantee or standby number.
	Identification *Max35Text `xml:"Id"`

	// Party that issues the undertaking.
	Issuer *PartyIdentification43 `xml:"Issr"`

	// Unique and unambiguous identifier assigned by the beneficiary to the undertaking.
	BeneficiaryReferenceNumber *Max35Text `xml:"BnfcryRefNb,omitempty"`
}

func (u *Undertaking6) SetIdentification(value string) {
	u.Identification = (*Max35Text)(&value)
}

func (u *Undertaking6) AddIssuer() *PartyIdentification43 {
	u.Issuer = new(PartyIdentification43)
	return u.Issuer
}

func (u *Undertaking6) SetBeneficiaryReferenceNumber(value string) {
	u.BeneficiaryReferenceNumber = (*Max35Text)(&value)
}
