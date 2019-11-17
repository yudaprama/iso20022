package model

// Information about an undertaking.
type Undertaking8 struct {

	// Unique and unambiguous identifier assigned by the issuer to the undertaking, for example the guarantee or standby number.
	Identification *Max35Text `xml:"Id"`

	// Party that issues the undertaking.
	Issuer *PartyIdentification43 `xml:"Issr"`

	// Unique and unambiguous identifier assigned by the applicant to the undertaking.
	ApplicantReferenceNumber *Max35Text `xml:"ApplcntRefNb,omitempty"`

	// Unique and unambiguous identifier assigned by the beneficiary to the undertaking.
	BeneficiaryReferenceNumber *Max35Text `xml:"BnfcryRefNb,omitempty"`
}

func (u *Undertaking8) SetIdentification(value string) {
	u.Identification = (*Max35Text)(&value)
}

func (u *Undertaking8) AddIssuer() *PartyIdentification43 {
	u.Issuer = new(PartyIdentification43)
	return u.Issuer
}

func (u *Undertaking8) SetApplicantReferenceNumber(value string) {
	u.ApplicantReferenceNumber = (*Max35Text)(&value)
}

func (u *Undertaking8) SetBeneficiaryReferenceNumber(value string) {
	u.BeneficiaryReferenceNumber = (*Max35Text)(&value)
}
