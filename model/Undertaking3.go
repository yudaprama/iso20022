package model

// Independent undertaking, such as a demand guarantee or standby letter of credit, that provides financial assurance, to be honoured on the presentation of documents that comply with its terms and conditions.
type Undertaking3 struct {

	// Unique and unambiguous identifier assigned by the issuer to the undertaking, for example the guarantee or standby number.
	Identification *Max35Text `xml:"Id"`

	// Name of undertaking such as, demand guarantee, standby letter of credit.
	Name *UndertakingIssuanceName1Code `xml:"Nm"`

	// Type of undertaking, for example, performance, payment.
	Type *UndertakingType1Choice `xml:"Tp,omitempty"`

	// Type of the undertaking issuance.
	IssuanceType *IssuanceType1Code `xml:"IssncTp"`

	// Party named in the undertaking as the “applicant”.
	Applicant []*PartyIdentification43 `xml:"Applcnt,omitempty"`

	// Party that issues the undertaking (or counter-undertaking).
	Issuer *PartyIdentification43 `xml:"Issr"`

	// Party in whose favour the undertaking (or counter-undertaking) is issued.
	Beneficiary []*PartyIdentification43 `xml:"Bnfcry"`

	// Date on which the undertaking is issued.
	DateOfIssuance *ISODate `xml:"DtOfIssnc"`

	// Location which is to be regarded as the place from which the undertaking is issued.
	PlaceOfIssue *PostalAddress12 `xml:"PlcOfIsse,omitempty"`

	// Party asked to advise the undertaking to the beneficiary or to another advising party at the request of the issuer.
	AdvisingParty *PartyIdentification43 `xml:"AdvsgPty,omitempty"`

	// Additional party asked to advise the undertaking.
	SecondAdvisingParty *PartyIdentification43 `xml:"ScndAdvsgPty,omitempty"`

	// Details related to the amount of the undertaking.
	UndertakingAmount *UndertakingAmount1 `xml:"UdrtkgAmt"`

	// Details related to the expiry of the undertaking.
	ExpiryDetails *ExpiryDetails1 `xml:"XpryDtls"`

	// Indicates whether or not the advising bank (confirmer) is requested to add its confirmation to the undertaking.
	ConfirmationIndicator *YesNoIndicator `xml:"ConfInd,omitempty"`

	// Indicates the type of party requested to add its confirmation to the undertaking.
	ConfirmationPartyType *ExternalTypeOfParty1Code `xml:"ConfPtyTp,omitempty"`

	// Party, in addition to the other parties specified in the undertaking, that is also related to the undertaking .
	AdditionalParty []*PartyAndType1 `xml:"AddtlPty,omitempty"`

	// Rules and laws governing the undertaking.
	GovernanceRulesAndLaw *GovernanceRules1 `xml:"GovncRulesAndLaw"`

	// Details of the underlying transaction for which the undertaking is issued.
	UnderlyingTransaction []*UnderlyingTradeTransaction1 `xml:"UndrlygTx,omitempty"`

	// Presentation details related to the undertaking.
	PresentationDetails *Presentation1 `xml:"PresntnDtls,omitempty"`

	// Terms and conditions of the undertaking.
	UndertakingTermsAndConditions []*Narrative1 `xml:"UdrtkgTermsAndConds"`

	// Indicates that multiple demands are not permitted.
	MultipleDemandIndicator *YesNoIndicator `xml:"MltplDmndInd,omitempty"`

	// Indicates that partial demands/drawings are not permitted.
	PartialDemandIndicator *YesNoIndicator `xml:"PrtlDmndInd,omitempty"`

	// Indicates whether the applicant/obligor or beneficiary is responsible for payment of the confirmation charges.
	ConfirmationChargesPayableBy *ExternalTypeOfParty1Code `xml:"ConfChrgsPyblBy,omitempty"`

	// Indicates whether the applicant/obligor or beneficiary is responsible for payment of the transfer charges.
	TransferChargesPayableBy *ExternalTypeOfParty1Code `xml:"TrfChrgsPyblBy,omitempty"`

	// Details related to a variation in amount that is automatically applied.
	AutomaticAmountVariation []*AutomaticVariation1 `xml:"AutomtcAmtVartn,omitempty"`

	// Details of the communication channel.
	DeliveryChannel *CommunicationChannel1 `xml:"DlvryChanl,omitempty"`

	// Indicates whether the undertaking is transferable.
	TransferIndicator *YesNoIndicator `xml:"TrfInd,omitempty"`

	// Document or template enclosed in the undertaking directly related to the issued undertaking.
	EnclosedFile []*Document9 `xml:"NclsdFile,omitempty"`

	// Additional information related to the undertaking.
	AdditionalInformation []*Max2000Text `xml:"AddtlInf,omitempty"`

	// Details of the local or ancillary undertaking requested to be issued by a local or other issuing institution.
	RequestedLocalUndertaking *Undertaking4 `xml:"ReqdLclUdrtkg,omitempty"`
}

func (u *Undertaking3) SetIdentification(value string) {
	u.Identification = (*Max35Text)(&value)
}

func (u *Undertaking3) SetName(value string) {
	u.Name = (*UndertakingIssuanceName1Code)(&value)
}

func (u *Undertaking3) AddType() *UndertakingType1Choice {
	u.Type = new(UndertakingType1Choice)
	return u.Type
}

func (u *Undertaking3) SetIssuanceType(value string) {
	u.IssuanceType = (*IssuanceType1Code)(&value)
}

func (u *Undertaking3) AddApplicant() *PartyIdentification43 {
	newValue := new(PartyIdentification43)
	u.Applicant = append(u.Applicant, newValue)
	return newValue
}

func (u *Undertaking3) AddIssuer() *PartyIdentification43 {
	u.Issuer = new(PartyIdentification43)
	return u.Issuer
}

func (u *Undertaking3) AddBeneficiary() *PartyIdentification43 {
	newValue := new(PartyIdentification43)
	u.Beneficiary = append(u.Beneficiary, newValue)
	return newValue
}

func (u *Undertaking3) SetDateOfIssuance(value string) {
	u.DateOfIssuance = (*ISODate)(&value)
}

func (u *Undertaking3) AddPlaceOfIssue() *PostalAddress12 {
	u.PlaceOfIssue = new(PostalAddress12)
	return u.PlaceOfIssue
}

func (u *Undertaking3) AddAdvisingParty() *PartyIdentification43 {
	u.AdvisingParty = new(PartyIdentification43)
	return u.AdvisingParty
}

func (u *Undertaking3) AddSecondAdvisingParty() *PartyIdentification43 {
	u.SecondAdvisingParty = new(PartyIdentification43)
	return u.SecondAdvisingParty
}

func (u *Undertaking3) AddUndertakingAmount() *UndertakingAmount1 {
	u.UndertakingAmount = new(UndertakingAmount1)
	return u.UndertakingAmount
}

func (u *Undertaking3) AddExpiryDetails() *ExpiryDetails1 {
	u.ExpiryDetails = new(ExpiryDetails1)
	return u.ExpiryDetails
}

func (u *Undertaking3) SetConfirmationIndicator(value string) {
	u.ConfirmationIndicator = (*YesNoIndicator)(&value)
}

func (u *Undertaking3) SetConfirmationPartyType(value string) {
	u.ConfirmationPartyType = (*ExternalTypeOfParty1Code)(&value)
}

func (u *Undertaking3) AddAdditionalParty() *PartyAndType1 {
	newValue := new(PartyAndType1)
	u.AdditionalParty = append(u.AdditionalParty, newValue)
	return newValue
}

func (u *Undertaking3) AddGovernanceRulesAndLaw() *GovernanceRules1 {
	u.GovernanceRulesAndLaw = new(GovernanceRules1)
	return u.GovernanceRulesAndLaw
}

func (u *Undertaking3) AddUnderlyingTransaction() *UnderlyingTradeTransaction1 {
	newValue := new(UnderlyingTradeTransaction1)
	u.UnderlyingTransaction = append(u.UnderlyingTransaction, newValue)
	return newValue
}

func (u *Undertaking3) AddPresentationDetails() *Presentation1 {
	u.PresentationDetails = new(Presentation1)
	return u.PresentationDetails
}

func (u *Undertaking3) AddUndertakingTermsAndConditions() *Narrative1 {
	newValue := new(Narrative1)
	u.UndertakingTermsAndConditions = append(u.UndertakingTermsAndConditions, newValue)
	return newValue
}

func (u *Undertaking3) SetMultipleDemandIndicator(value string) {
	u.MultipleDemandIndicator = (*YesNoIndicator)(&value)
}

func (u *Undertaking3) SetPartialDemandIndicator(value string) {
	u.PartialDemandIndicator = (*YesNoIndicator)(&value)
}

func (u *Undertaking3) SetConfirmationChargesPayableBy(value string) {
	u.ConfirmationChargesPayableBy = (*ExternalTypeOfParty1Code)(&value)
}

func (u *Undertaking3) SetTransferChargesPayableBy(value string) {
	u.TransferChargesPayableBy = (*ExternalTypeOfParty1Code)(&value)
}

func (u *Undertaking3) AddAutomaticAmountVariation() *AutomaticVariation1 {
	newValue := new(AutomaticVariation1)
	u.AutomaticAmountVariation = append(u.AutomaticAmountVariation, newValue)
	return newValue
}

func (u *Undertaking3) AddDeliveryChannel() *CommunicationChannel1 {
	u.DeliveryChannel = new(CommunicationChannel1)
	return u.DeliveryChannel
}

func (u *Undertaking3) SetTransferIndicator(value string) {
	u.TransferIndicator = (*YesNoIndicator)(&value)
}

func (u *Undertaking3) AddEnclosedFile() *Document9 {
	newValue := new(Document9)
	u.EnclosedFile = append(u.EnclosedFile, newValue)
	return newValue
}

func (u *Undertaking3) AddAdditionalInformation(value string) {
	u.AdditionalInformation = append(u.AdditionalInformation, (*Max2000Text)(&value))
}

func (u *Undertaking3) AddRequestedLocalUndertaking() *Undertaking4 {
	u.RequestedLocalUndertaking = new(Undertaking4)
	return u.RequestedLocalUndertaking
}
