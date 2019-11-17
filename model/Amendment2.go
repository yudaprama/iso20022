package model

// Details of the amendment.
type Amendment2 struct {

	// Contents of the related Undertaking Amendment message.
	UndertakingAmendmentMessage *UndertakingAmendmentMessage1 `xml:"UdrtkgAmdmntMsg"`

	// Additional information related to the first advising party.
	FirstAdvisingPartyAdditionalInformation *AdvisingPartyAdditionalInformation1 `xml:"FrstAdvsgPtyAddtlInf,omitempty"`

	// Additional information related to the second advising party.
	SecondAdvisingPartyAdditionalInformation *AdvisingPartyAdditionalInformation1 `xml:"ScndAdvsgPtyAddtlInf,omitempty"`

	// Details concerning the confirmation of the proposed amendment.
	ConfirmationDetails *UndertakingConfirmation1 `xml:"ConfDtls,omitempty"`

	// Digital signature of the party providing additional undertaking amendment advice details.
	DigitalSignature []*PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`
}

func (a *Amendment2) AddUndertakingAmendmentMessage() *UndertakingAmendmentMessage1 {
	a.UndertakingAmendmentMessage = new(UndertakingAmendmentMessage1)
	return a.UndertakingAmendmentMessage
}

func (a *Amendment2) AddFirstAdvisingPartyAdditionalInformation() *AdvisingPartyAdditionalInformation1 {
	a.FirstAdvisingPartyAdditionalInformation = new(AdvisingPartyAdditionalInformation1)
	return a.FirstAdvisingPartyAdditionalInformation
}

func (a *Amendment2) AddSecondAdvisingPartyAdditionalInformation() *AdvisingPartyAdditionalInformation1 {
	a.SecondAdvisingPartyAdditionalInformation = new(AdvisingPartyAdditionalInformation1)
	return a.SecondAdvisingPartyAdditionalInformation
}

func (a *Amendment2) AddConfirmationDetails() *UndertakingConfirmation1 {
	a.ConfirmationDetails = new(UndertakingConfirmation1)
	return a.ConfirmationDetails
}

func (a *Amendment2) AddDigitalSignature() *PartyAndSignature2 {
	newValue := new(PartyAndSignature2)
	a.DigitalSignature = append(a.DigitalSignature, newValue)
	return newValue
}
