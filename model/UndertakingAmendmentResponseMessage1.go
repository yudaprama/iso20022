package model

// Contents of an UndertakingAmendmentResponse message.
type UndertakingAmendmentResponseMessage1 struct {

	// Details of the proposed amendment response.
	UndertakingAmendmentResponseDetails *Amendment7 `xml:"UdrtkgAmdmntRspnDtls"`

	// Digital signature of the response.
	DigitalSignature *PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`
}

func (u *UndertakingAmendmentResponseMessage1) AddUndertakingAmendmentResponseDetails() *Amendment7 {
	u.UndertakingAmendmentResponseDetails = new(Amendment7)
	return u.UndertakingAmendmentResponseDetails
}

func (u *UndertakingAmendmentResponseMessage1) AddDigitalSignature() *PartyAndSignature2 {
	u.DigitalSignature = new(PartyAndSignature2)
	return u.DigitalSignature
}
