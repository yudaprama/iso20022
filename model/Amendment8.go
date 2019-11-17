package model

// Amendment identification.
type Amendment8 struct {

	// Unique and unambiguous identifier assigned by the issuer to the undertaking, for example the guarantee or standby number.
	Identification *Max35Text `xml:"Id"`

	// Sequence number assigned by the issuer to each amendment of the undertaking.
	SequenceNumber *Max4AlphaNumericText `xml:"SeqNb"`

	// Unique and unambiguous identifier assigned by the beneficiary to the undertaking.
	BeneficiaryReferenceNumber *Max35Text `xml:"BnfcryRefNb,omitempty"`

	// Party that issues the undertaking.
	Issuer *PartyIdentification43 `xml:"Issr"`
}

func (a *Amendment8) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *Amendment8) SetSequenceNumber(value string) {
	a.SequenceNumber = (*Max4AlphaNumericText)(&value)
}

func (a *Amendment8) SetBeneficiaryReferenceNumber(value string) {
	a.BeneficiaryReferenceNumber = (*Max35Text)(&value)
}

func (a *Amendment8) AddIssuer() *PartyIdentification43 {
	a.Issuer = new(PartyIdentification43)
	return a.Issuer
}
