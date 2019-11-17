package model

// Details of the amendment.
type Amendment1 struct {

	// Sequence number assigned by the issuer to each proposed amendment of the undertaking.
	SequenceNumber *Max4AlphaNumericText `xml:"SeqNb"`

	// Date on which the proposed amendment is issued.
	DateOfIssuance *ISODate `xml:"DtOfIssnc"`

	// Identification of the undertaking.
	UndertakingIdentification *Undertaking7 `xml:"UdrtkgId"`

	// Party asked to advise the proposed amendment to the beneficiary or to another advising party at the request of the issuer.
	AdvisingParty *PartyIdentification43 `xml:"AdvsgPty,omitempty"`

	// Additional party asked to advise the proposed amendment.
	SecondAdvisingParty *PartyIdentification43 `xml:"ScndAdvsgPty,omitempty"`

	// Details concerning the requested termination of the undertaking.
	TerminationDetails *UndertakingTermination3 `xml:"TermntnDtls,omitempty"`

	// Requested increase or decrease to the amount of for the undertaking.
	UndertakingAmountAdjustment *UndertakingAmount2 `xml:"UdrtkgAmtAdjstmnt,omitempty"`

	// Requested new expiry terms for the undertaking.
	NewExpiryDetails *ExpiryDetails1 `xml:"NewXpryDtls,omitempty"`

	// Requested new beneficiary of the undertaking.
	NewBeneficiary *PartyIdentification43 `xml:"NewBnfcry,omitempty"`

	// Requested new terms and conditions of the undertaking.
	NewUndertakingTermsAndConditions []*Narrative1 `xml:"NewUdrtkgTermsAndConds,omitempty"`

	// Amendment details related to the local undertaking.
	LocalUndertaking *Undertaking11 `xml:"LclUdrtkg,omitempty"`

	// Indicates whether or not consent is requested from the beneficiary.
	BeneficiaryConsentRequestIndicator *YesNoIndicator `xml:"BnfcryCnsntReqInd,omitempty"`

	// Communication channel for delivery of the proposed amendment.
	DeliveryChannel *CommunicationChannel1 `xml:"DlvryChanl,omitempty"`

	// Document or template enclosed in the proposed amendment.
	EnclosedFile []*Document9 `xml:"NclsdFile,omitempty"`

	// Additional information related to the proposed amendment.
	AdditionalInformation []*Max2000Text `xml:"AddtlInf,omitempty"`
}

func (a *Amendment1) SetSequenceNumber(value string) {
	a.SequenceNumber = (*Max4AlphaNumericText)(&value)
}

func (a *Amendment1) SetDateOfIssuance(value string) {
	a.DateOfIssuance = (*ISODate)(&value)
}

func (a *Amendment1) AddUndertakingIdentification() *Undertaking7 {
	a.UndertakingIdentification = new(Undertaking7)
	return a.UndertakingIdentification
}

func (a *Amendment1) AddAdvisingParty() *PartyIdentification43 {
	a.AdvisingParty = new(PartyIdentification43)
	return a.AdvisingParty
}

func (a *Amendment1) AddSecondAdvisingParty() *PartyIdentification43 {
	a.SecondAdvisingParty = new(PartyIdentification43)
	return a.SecondAdvisingParty
}

func (a *Amendment1) AddTerminationDetails() *UndertakingTermination3 {
	a.TerminationDetails = new(UndertakingTermination3)
	return a.TerminationDetails
}

func (a *Amendment1) AddUndertakingAmountAdjustment() *UndertakingAmount2 {
	a.UndertakingAmountAdjustment = new(UndertakingAmount2)
	return a.UndertakingAmountAdjustment
}

func (a *Amendment1) AddNewExpiryDetails() *ExpiryDetails1 {
	a.NewExpiryDetails = new(ExpiryDetails1)
	return a.NewExpiryDetails
}

func (a *Amendment1) AddNewBeneficiary() *PartyIdentification43 {
	a.NewBeneficiary = new(PartyIdentification43)
	return a.NewBeneficiary
}

func (a *Amendment1) AddNewUndertakingTermsAndConditions() *Narrative1 {
	newValue := new(Narrative1)
	a.NewUndertakingTermsAndConditions = append(a.NewUndertakingTermsAndConditions, newValue)
	return newValue
}

func (a *Amendment1) AddLocalUndertaking() *Undertaking11 {
	a.LocalUndertaking = new(Undertaking11)
	return a.LocalUndertaking
}

func (a *Amendment1) SetBeneficiaryConsentRequestIndicator(value string) {
	a.BeneficiaryConsentRequestIndicator = (*YesNoIndicator)(&value)
}

func (a *Amendment1) AddDeliveryChannel() *CommunicationChannel1 {
	a.DeliveryChannel = new(CommunicationChannel1)
	return a.DeliveryChannel
}

func (a *Amendment1) AddEnclosedFile() *Document9 {
	newValue := new(Document9)
	a.EnclosedFile = append(a.EnclosedFile, newValue)
	return newValue
}

func (a *Amendment1) AddAdditionalInformation(value string) {
	a.AdditionalInformation = append(a.AdditionalInformation, (*Max2000Text)(&value))
}
