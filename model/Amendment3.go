package model

// Details of the amendent request.
type Amendment3 struct {

	// Unique and unambiguous identifier assigned by the applicant to the undertaking amendment request.
	ApplicantRequestNumber *Max35Text `xml:"ApplcntReqNb"`

	// Identification of the undertaking.
	UndertakingIdentification *Undertaking9 `xml:"UdrtkgId"`

	// Party requesting the issuance of the amendment.
	Applicant *PartyIdentification43 `xml:"Applcnt"`

	// Details concerning the requested termination of the undertaking.
	TerminationDetails *UndertakingTermination3 `xml:"TermntnDtls,omitempty"`

	// Indication of the amount of increase or decrease to the undertaking amount.
	IncreaseDecreaseAmount *UndertakingAmount2 `xml:"IncrDcrAmt,omitempty"`

	// Requested new expiry terms for the undertaking.
	NewExpiryDetails *ExpiryDetails2 `xml:"NewXpryDtls,omitempty"`

	// Requested new beneficiary of the undertaking.
	NewBeneficiary *Beneficiary1 `xml:"NewBnfcry,omitempty"`

	// Requested new terms and conditions of the undertaking.
	NewUndertakingTermsAndConditions []*Narrative1 `xml:"NewUdrtkgTermsAndConds,omitempty"`

	// Amendment details related to the counter-undertaking.
	CounterUndertaking *Undertaking10 `xml:"CntrUdrtkg,omitempty"`

	// Communication channel for delivery of the amendment.
	DeliveryChannel *CommunicationChannel1 `xml:"DlvryChanl,omitempty"`

	// Document or template enclosed in the request.
	EnclosedFile []*Document9 `xml:"NclsdFile,omitempty"`

	// Additional information related to the request.
	AdditionalInformation []*Max2000Text `xml:"AddtlInf,omitempty"`
}

func (a *Amendment3) SetApplicantRequestNumber(value string) {
	a.ApplicantRequestNumber = (*Max35Text)(&value)
}

func (a *Amendment3) AddUndertakingIdentification() *Undertaking9 {
	a.UndertakingIdentification = new(Undertaking9)
	return a.UndertakingIdentification
}

func (a *Amendment3) AddApplicant() *PartyIdentification43 {
	a.Applicant = new(PartyIdentification43)
	return a.Applicant
}

func (a *Amendment3) AddTerminationDetails() *UndertakingTermination3 {
	a.TerminationDetails = new(UndertakingTermination3)
	return a.TerminationDetails
}

func (a *Amendment3) AddIncreaseDecreaseAmount() *UndertakingAmount2 {
	a.IncreaseDecreaseAmount = new(UndertakingAmount2)
	return a.IncreaseDecreaseAmount
}

func (a *Amendment3) AddNewExpiryDetails() *ExpiryDetails2 {
	a.NewExpiryDetails = new(ExpiryDetails2)
	return a.NewExpiryDetails
}

func (a *Amendment3) AddNewBeneficiary() *Beneficiary1 {
	a.NewBeneficiary = new(Beneficiary1)
	return a.NewBeneficiary
}

func (a *Amendment3) AddNewUndertakingTermsAndConditions() *Narrative1 {
	newValue := new(Narrative1)
	a.NewUndertakingTermsAndConditions = append(a.NewUndertakingTermsAndConditions, newValue)
	return newValue
}

func (a *Amendment3) AddCounterUndertaking() *Undertaking10 {
	a.CounterUndertaking = new(Undertaking10)
	return a.CounterUndertaking
}

func (a *Amendment3) AddDeliveryChannel() *CommunicationChannel1 {
	a.DeliveryChannel = new(CommunicationChannel1)
	return a.DeliveryChannel
}

func (a *Amendment3) AddEnclosedFile() *Document9 {
	newValue := new(Document9)
	a.EnclosedFile = append(a.EnclosedFile, newValue)
	return newValue
}

func (a *Amendment3) AddAdditionalInformation(value string) {
	a.AdditionalInformation = append(a.AdditionalInformation, (*Max2000Text)(&value))
}
