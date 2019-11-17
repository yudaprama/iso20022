package model

// Contents of an Undertaking Amendment message.
type UndertakingAmendmentMessage1 struct {

	// Details related to the proposed undertaking amendment.
	UndertakingAmendmentDetails *Amendment1 `xml:"UdrtkgAmdmntDtls"`

	// Digital signature of the proposed amendment.
	DigitalSignature *PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`
}

func (u *UndertakingAmendmentMessage1) AddUndertakingAmendmentDetails() *Amendment1 {
	u.UndertakingAmendmentDetails = new(Amendment1)
	return u.UndertakingAmendmentDetails
}

func (u *UndertakingAmendmentMessage1) AddDigitalSignature() *PartyAndSignature2 {
	u.DigitalSignature = new(PartyAndSignature2)
	return u.DigitalSignature
}
