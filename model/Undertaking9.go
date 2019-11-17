package model

// Information about an undertaking.
type Undertaking9 struct {

	// Unique and unambiguous identifier assigned by the issuer to the undertaking, for example the guarantee or standby number.
	Identification *Max35Text `xml:"Id"`

	// Party that issues the undertaking.
	Issuer *PartyIdentification43 `xml:"Issr"`

	// Unique and unambiguous identifier assigned by the applicant to the undertaking.
	ApplicantReferenceNumber *Max35Text `xml:"ApplcntRefNb,omitempty"`
}

func (u *Undertaking9) SetIdentification(value string) {
	u.Identification = (*Max35Text)(&value)
}

func (u *Undertaking9) AddIssuer() *PartyIdentification43 {
	u.Issuer = new(PartyIdentification43)
	return u.Issuer
}

func (u *Undertaking9) SetApplicantReferenceNumber(value string) {
	u.ApplicantReferenceNumber = (*Max35Text)(&value)
}
