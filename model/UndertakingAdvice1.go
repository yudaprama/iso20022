package model

// Details of the advice for the issuance of an undertaking.
type UndertakingAdvice1 struct {

	// Contents of the related UndertakingIssuance message.
	UndertakingIssuanceMessage *UndertakingIssuanceMessage `xml:"UdrtkgIssncMsg"`

	// Additional information related to the first advising party.
	FirstAdvisingPartyAdditionalInformation *AdvisingPartyAdditionalInformation1 `xml:"FrstAdvsgPtyAddtlInf,omitempty"`

	// Additional information related to the second advising party.
	SecondAdvisingPartyAdditionalInformation *AdvisingPartyAdditionalInformation1 `xml:"ScndAdvsgPtyAddtlInf,omitempty"`

	// Details related to the confirmation of the undertaking.
	ConfirmationDetails *UndertakingConfirmation1 `xml:"ConfDtls,omitempty"`

	// Digital signature of the party providing additional undertaking advice details.
	DigitalSignature []*PartyAndSignature2 `xml:"DgtlSgntr,omitempty"`
}

func (u *UndertakingAdvice1) AddUndertakingIssuanceMessage() *UndertakingIssuanceMessage {
	u.UndertakingIssuanceMessage = new(UndertakingIssuanceMessage)
	return u.UndertakingIssuanceMessage
}

func (u *UndertakingAdvice1) AddFirstAdvisingPartyAdditionalInformation() *AdvisingPartyAdditionalInformation1 {
	u.FirstAdvisingPartyAdditionalInformation = new(AdvisingPartyAdditionalInformation1)
	return u.FirstAdvisingPartyAdditionalInformation
}

func (u *UndertakingAdvice1) AddSecondAdvisingPartyAdditionalInformation() *AdvisingPartyAdditionalInformation1 {
	u.SecondAdvisingPartyAdditionalInformation = new(AdvisingPartyAdditionalInformation1)
	return u.SecondAdvisingPartyAdditionalInformation
}

func (u *UndertakingAdvice1) AddConfirmationDetails() *UndertakingConfirmation1 {
	u.ConfirmationDetails = new(UndertakingConfirmation1)
	return u.ConfirmationDetails
}

func (u *UndertakingAdvice1) AddDigitalSignature() *PartyAndSignature2 {
	newValue := new(PartyAndSignature2)
	u.DigitalSignature = append(u.DigitalSignature, newValue)
	return newValue
}
