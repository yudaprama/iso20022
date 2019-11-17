package model

// Independent undertaking, such as a demand guarantee or standby letter of credit, that provides financial assurance, to be honoured on the presentation of documents that comply with its terms and conditions.
type Undertaking1 struct {

	// Unique and unambiguous identifier assigned by the applicant to the undertaking.
	ApplicantReferenceNumber *Max35Text `xml:"ApplcntRefNb"`

	// Brief description of the purpose of the undertaking. Provided as information for the issuer reference.
	Purpose *Max350Text `xml:"Purp"`

	// Undertaking name.
	Name *UndertakingName1Code `xml:"Nm"`

	// Type of undertaking, for example, performance, payment.
	Type *UndertakingType1Choice `xml:"Tp"`

	// Party obligated to reimburse the issuer.
	Obligor *PartyIdentification43 `xml:"Oblgr"`

	// Party to be named in the undertaking as the “applicant” when different from the obligor.
	Applicant []*PartyIdentification43 `xml:"Applcnt,omitempty"`

	// Party that issues the undertaking (or counter-undertaking).
	Issuer *PartyIdentification43 `xml:"Issr"`

	// Ultimate party in whose favour the undertaking is to be issued.
	Beneficiary []*PartyIdentification43 `xml:"Bnfcry"`

	// Party asked to advise the undertaking to the beneficiary or to another advising party.
	AdvisingParty *PartyIdentification43 `xml:"AdvsgPty,omitempty"`

	// Additional party asked to advise the undertaking.
	SecondAdvisingParty *PartyIdentification43 `xml:"ScndAdvsgPty,omitempty"`

	// Party that adds its confirmation to the undertaking. For further clarification, reference the applicable rules to which the undertaking is subject.
	Confirmer *PartyIdentification43 `xml:"Cnfrmr,omitempty"`

	// Indicates whether the advising bank (confirmer) is requested to add its confirmation to the undertaking. The absence of this element indicates that the advising bank (confirmer) is not requested to add its confirmation to the undertaking.
	ConfirmationIndicator *YesNoIndicator `xml:"ConfInd,omitempty"`

	// Indicates whether the undertaking is a local or ancillary undertaking to be issued under a counter-undertaking.
	CounterUndertakingIndicator *YesNoIndicator `xml:"CntrUdrtkgInd"`

	// Details related to the counter undertaking.
	CounterUndertaking *Undertaking2 `xml:"CntrUdrtkg,omitempty"`

	// Details related to the amount of the undertaking.
	UndertakingAmount *UndertakingAmount1 `xml:"UdrtkgAmt"`

	// Details related to the expiry terms of the undertaking.
	ExpiryDetails *ExpiryDetails2 `xml:"XpryDtls"`

	// Party, in addition to the other parties specified in the undertaking, that is also related to the undertaking.
	AdditionalParty []*PartyAndType1 `xml:"AddtlPty,omitempty"`

	// Rules and laws governing the undertaking.
	GovernanceRulesAndLaw *GovernanceRules1 `xml:"GovncRulesAndLaw"`

	// Details of the underlying transaction for which the undertaking is issued.
	UnderlyingTransaction []*UnderlyingTradeTransaction1 `xml:"UndrlygTx,omitempty"`

	// Presentation details related to the undertaking.
	PresentationDetails *Presentation4 `xml:"PresntnDtls,omitempty"`

	// Wording details and text for the undertaking.
	UndertakingWording *UndertakingWording1 `xml:"UdrtkgWrdg"`

	// Indicates whether multiple demands are permitted.
	MultipleDemandIndicator *YesNoIndicator `xml:"MltplDmndInd,omitempty"`

	// Indicates whether partial demands/drawings are permitted.
	PartialDemandIndicator *YesNoIndicator `xml:"PrtlDmndInd,omitempty"`

	// Indicates whether the undertaking is transferable.
	TransferIndicator *YesNoIndicator `xml:"TrfInd,omitempty"`

	// Indicates whether the applicant/obligor or beneficiary is responsible for payment of the transfer charges.
	TransferChargesPayableBy *ExternalTypeOfParty1Code `xml:"TrfChrgsPyblBy,omitempty"`

	// Indicates whether the applicant/obligor or beneficiary is responsible for payment of the confirmation charges.
	ConfirmationChargesPayableBy *ExternalTypeOfParty1Code `xml:"ConfChrgsPyblBy,omitempty"`

	// Details related to a variation in amount that is automatically applied.
	AutomaticAmountVariation []*AutomaticVariation1 `xml:"AutomtcAmtVartn,omitempty"`

	// Details of the communication channel.
	DeliveryChannel *CommunicationChannel1 `xml:"DlvryChanl"`

	// Account nominated by the obligor to record the liability amount of the undertaking.
	ObligorLiabilityAccount *CashAccount28 `xml:"OblgrLbltyAcct,omitempty"`

	// Account nominated by the obligor for the settlement of charges related to the undertaking.
	ObligorChargeAccount *CashAccount28 `xml:"OblgrChrgAcct,omitempty"`

	// Account nominated by the obligor for the settlement of the amount claimed under the undertaking.
	ObligorSettlementAccount *CashAccount28 `xml:"OblgrSttlmAcct,omitempty"`

	// Document or template enclosed in the undertaking directly related to the undertaking to be issued.
	EnclosedFile []*Document9 `xml:"NclsdFile,omitempty"`

	// Additional information related to the application for an undertaking.
	AdditionalApplicationInformation []*Max2000Text `xml:"AddtlApplInf,omitempty"`
}

func (u *Undertaking1) SetApplicantReferenceNumber(value string) {
	u.ApplicantReferenceNumber = (*Max35Text)(&value)
}

func (u *Undertaking1) SetPurpose(value string) {
	u.Purpose = (*Max350Text)(&value)
}

func (u *Undertaking1) SetName(value string) {
	u.Name = (*UndertakingName1Code)(&value)
}

func (u *Undertaking1) AddType() *UndertakingType1Choice {
	u.Type = new(UndertakingType1Choice)
	return u.Type
}

func (u *Undertaking1) AddObligor() *PartyIdentification43 {
	u.Obligor = new(PartyIdentification43)
	return u.Obligor
}

func (u *Undertaking1) AddApplicant() *PartyIdentification43 {
	newValue := new(PartyIdentification43)
	u.Applicant = append(u.Applicant, newValue)
	return newValue
}

func (u *Undertaking1) AddIssuer() *PartyIdentification43 {
	u.Issuer = new(PartyIdentification43)
	return u.Issuer
}

func (u *Undertaking1) AddBeneficiary() *PartyIdentification43 {
	newValue := new(PartyIdentification43)
	u.Beneficiary = append(u.Beneficiary, newValue)
	return newValue
}

func (u *Undertaking1) AddAdvisingParty() *PartyIdentification43 {
	u.AdvisingParty = new(PartyIdentification43)
	return u.AdvisingParty
}

func (u *Undertaking1) AddSecondAdvisingParty() *PartyIdentification43 {
	u.SecondAdvisingParty = new(PartyIdentification43)
	return u.SecondAdvisingParty
}

func (u *Undertaking1) AddConfirmer() *PartyIdentification43 {
	u.Confirmer = new(PartyIdentification43)
	return u.Confirmer
}

func (u *Undertaking1) SetConfirmationIndicator(value string) {
	u.ConfirmationIndicator = (*YesNoIndicator)(&value)
}

func (u *Undertaking1) SetCounterUndertakingIndicator(value string) {
	u.CounterUndertakingIndicator = (*YesNoIndicator)(&value)
}

func (u *Undertaking1) AddCounterUndertaking() *Undertaking2 {
	u.CounterUndertaking = new(Undertaking2)
	return u.CounterUndertaking
}

func (u *Undertaking1) AddUndertakingAmount() *UndertakingAmount1 {
	u.UndertakingAmount = new(UndertakingAmount1)
	return u.UndertakingAmount
}

func (u *Undertaking1) AddExpiryDetails() *ExpiryDetails2 {
	u.ExpiryDetails = new(ExpiryDetails2)
	return u.ExpiryDetails
}

func (u *Undertaking1) AddAdditionalParty() *PartyAndType1 {
	newValue := new(PartyAndType1)
	u.AdditionalParty = append(u.AdditionalParty, newValue)
	return newValue
}

func (u *Undertaking1) AddGovernanceRulesAndLaw() *GovernanceRules1 {
	u.GovernanceRulesAndLaw = new(GovernanceRules1)
	return u.GovernanceRulesAndLaw
}

func (u *Undertaking1) AddUnderlyingTransaction() *UnderlyingTradeTransaction1 {
	newValue := new(UnderlyingTradeTransaction1)
	u.UnderlyingTransaction = append(u.UnderlyingTransaction, newValue)
	return newValue
}

func (u *Undertaking1) AddPresentationDetails() *Presentation4 {
	u.PresentationDetails = new(Presentation4)
	return u.PresentationDetails
}

func (u *Undertaking1) AddUndertakingWording() *UndertakingWording1 {
	u.UndertakingWording = new(UndertakingWording1)
	return u.UndertakingWording
}

func (u *Undertaking1) SetMultipleDemandIndicator(value string) {
	u.MultipleDemandIndicator = (*YesNoIndicator)(&value)
}

func (u *Undertaking1) SetPartialDemandIndicator(value string) {
	u.PartialDemandIndicator = (*YesNoIndicator)(&value)
}

func (u *Undertaking1) SetTransferIndicator(value string) {
	u.TransferIndicator = (*YesNoIndicator)(&value)
}

func (u *Undertaking1) SetTransferChargesPayableBy(value string) {
	u.TransferChargesPayableBy = (*ExternalTypeOfParty1Code)(&value)
}

func (u *Undertaking1) SetConfirmationChargesPayableBy(value string) {
	u.ConfirmationChargesPayableBy = (*ExternalTypeOfParty1Code)(&value)
}

func (u *Undertaking1) AddAutomaticAmountVariation() *AutomaticVariation1 {
	newValue := new(AutomaticVariation1)
	u.AutomaticAmountVariation = append(u.AutomaticAmountVariation, newValue)
	return newValue
}

func (u *Undertaking1) AddDeliveryChannel() *CommunicationChannel1 {
	u.DeliveryChannel = new(CommunicationChannel1)
	return u.DeliveryChannel
}

func (u *Undertaking1) AddObligorLiabilityAccount() *CashAccount28 {
	u.ObligorLiabilityAccount = new(CashAccount28)
	return u.ObligorLiabilityAccount
}

func (u *Undertaking1) AddObligorChargeAccount() *CashAccount28 {
	u.ObligorChargeAccount = new(CashAccount28)
	return u.ObligorChargeAccount
}

func (u *Undertaking1) AddObligorSettlementAccount() *CashAccount28 {
	u.ObligorSettlementAccount = new(CashAccount28)
	return u.ObligorSettlementAccount
}

func (u *Undertaking1) AddEnclosedFile() *Document9 {
	newValue := new(Document9)
	u.EnclosedFile = append(u.EnclosedFile, newValue)
	return newValue
}

func (u *Undertaking1) AddAdditionalApplicationInformation(value string) {
	u.AdditionalApplicationInformation = append(u.AdditionalApplicationInformation, (*Max2000Text)(&value))
}
