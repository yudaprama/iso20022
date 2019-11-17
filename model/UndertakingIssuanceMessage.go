package model

// Contents of the related UndertakingIssuance message or of the related issuance document.
type UndertakingIssuanceMessage struct {

	// Independent undertaking, such as a demand guarantee or standby letter of credit, that provides financial assurance, to be collected on the presentation of documents that comply with its terms and conditions.
	UndertakingDetails *Undertaking3 `xml:"UdrtkgDtls"`

	// Digital signature of the issued undertaking.
	DigitalSignature *PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`
}

func (u *UndertakingIssuanceMessage) AddUndertakingDetails() *Undertaking3 {
	u.UndertakingDetails = new(Undertaking3)
	return u.UndertakingDetails
}

func (u *UndertakingIssuanceMessage) AddDigitalSignature() *PartyAndSignature2 {
	u.DigitalSignature = new(PartyAndSignature2)
	return u.DigitalSignature
}
