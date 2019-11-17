package model

// Information about an undertaking.
type Undertaking4 struct {

	// Name of the requested local undertaking such as, demand guarantee, standby letter of credit, surety.
	Name *UndertakingName1Code `xml:"Nm"`

	// Type of the requested local undertaking such as performance, payment.
	Type *ExternalUndertakingType1Code `xml:"Tp"`

	// Party requested to be named in the local undertaking as the party on whose behalf the undertaking is issued.
	Applicant []*PartyIdentification43 `xml:"Applcnt"`

	// Party in whose favour the requested local undertaking is to be issued.
	Beneficiary []*PartyIdentification43 `xml:"Bnfcry"`

	// Date on which the requested local undertaking is to be issued.
	DateOfIssuance *ISODate `xml:"DtOfIssnc,omitempty"`

	// Party asked to advise the requested local undertaking to the beneficiary or to another advising party at the request of the local undertaking issuer.
	AdvisingParty *PartyIdentification43 `xml:"AdvsgPty,omitempty"`

	// Additional party asked to advise the requested local undertaking.
	SecondAdvisingParty *PartyIdentification43 `xml:"ScndAdvsgPty,omitempty"`

	// Details related to the amount of the local undertaking.
	LocalUndertakingAmount *UndertakingAmount1 `xml:"LclUdrtkgAmt"`

	// Details related to the expiry of the requested local undertaking.
	ExpiryDetails *ExpiryDetails1 `xml:"XpryDtls"`

	// Indicates whether or not the advising bank (confirmer) is requested to add its confirmation to the undertaking.
	ConfirmationIndicator *YesNoIndicator `xml:"ConfInd,omitempty"`

	// Party, in addition to the other parties specified in the requested local undertaking, that is also related to the requested local undertaking.
	AdditionalParty []*PartyAndType1 `xml:"AddtlPty,omitempty"`

	// Rules and laws governing the requested local undertaking.
	GovernanceRulesAndLaw *GovernanceRules1 `xml:"GovncRulesAndLaw"`

	// Details of the underlying transaction for which the undertaking is issued.
	UnderlyingTransaction []*UnderlyingTradeTransaction1 `xml:"UndrlygTx,omitempty"`

	// Presentation details related to the undertaking.
	PresentationDetails *Presentation1 `xml:"PresntnDtls,omitempty"`

	// Wording details and text for the requested local undertaking.
	UndertakingWording *UndertakingWording1 `xml:"UdrtkgWrdg"`

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

	// Indicates whether the requested local undertaking is transferable.
	TransferIndicator *YesNoIndicator `xml:"TrfInd,omitempty"`

	// Additional information related to the requested local undertaking.
	AdditionalInformation []*Max2000Text `xml:"AddtlInf,omitempty"`
}

func (u *Undertaking4) SetName(value string) {
	u.Name = (*UndertakingName1Code)(&value)
}

func (u *Undertaking4) SetType(value string) {
	u.Type = (*ExternalUndertakingType1Code)(&value)
}

func (u *Undertaking4) AddApplicant() *PartyIdentification43 {
	newValue := new(PartyIdentification43)
	u.Applicant = append(u.Applicant, newValue)
	return newValue
}

func (u *Undertaking4) AddBeneficiary() *PartyIdentification43 {
	newValue := new(PartyIdentification43)
	u.Beneficiary = append(u.Beneficiary, newValue)
	return newValue
}

func (u *Undertaking4) SetDateOfIssuance(value string) {
	u.DateOfIssuance = (*ISODate)(&value)
}

func (u *Undertaking4) AddAdvisingParty() *PartyIdentification43 {
	u.AdvisingParty = new(PartyIdentification43)
	return u.AdvisingParty
}

func (u *Undertaking4) AddSecondAdvisingParty() *PartyIdentification43 {
	u.SecondAdvisingParty = new(PartyIdentification43)
	return u.SecondAdvisingParty
}

func (u *Undertaking4) AddLocalUndertakingAmount() *UndertakingAmount1 {
	u.LocalUndertakingAmount = new(UndertakingAmount1)
	return u.LocalUndertakingAmount
}

func (u *Undertaking4) AddExpiryDetails() *ExpiryDetails1 {
	u.ExpiryDetails = new(ExpiryDetails1)
	return u.ExpiryDetails
}

func (u *Undertaking4) SetConfirmationIndicator(value string) {
	u.ConfirmationIndicator = (*YesNoIndicator)(&value)
}

func (u *Undertaking4) AddAdditionalParty() *PartyAndType1 {
	newValue := new(PartyAndType1)
	u.AdditionalParty = append(u.AdditionalParty, newValue)
	return newValue
}

func (u *Undertaking4) AddGovernanceRulesAndLaw() *GovernanceRules1 {
	u.GovernanceRulesAndLaw = new(GovernanceRules1)
	return u.GovernanceRulesAndLaw
}

func (u *Undertaking4) AddUnderlyingTransaction() *UnderlyingTradeTransaction1 {
	newValue := new(UnderlyingTradeTransaction1)
	u.UnderlyingTransaction = append(u.UnderlyingTransaction, newValue)
	return newValue
}

func (u *Undertaking4) AddPresentationDetails() *Presentation1 {
	u.PresentationDetails = new(Presentation1)
	return u.PresentationDetails
}

func (u *Undertaking4) AddUndertakingWording() *UndertakingWording1 {
	u.UndertakingWording = new(UndertakingWording1)
	return u.UndertakingWording
}

func (u *Undertaking4) SetMultipleDemandIndicator(value string) {
	u.MultipleDemandIndicator = (*YesNoIndicator)(&value)
}

func (u *Undertaking4) SetPartialDemandIndicator(value string) {
	u.PartialDemandIndicator = (*YesNoIndicator)(&value)
}

func (u *Undertaking4) SetConfirmationChargesPayableBy(value string) {
	u.ConfirmationChargesPayableBy = (*ExternalTypeOfParty1Code)(&value)
}

func (u *Undertaking4) SetTransferChargesPayableBy(value string) {
	u.TransferChargesPayableBy = (*ExternalTypeOfParty1Code)(&value)
}

func (u *Undertaking4) AddAutomaticAmountVariation() *AutomaticVariation1 {
	newValue := new(AutomaticVariation1)
	u.AutomaticAmountVariation = append(u.AutomaticAmountVariation, newValue)
	return newValue
}

func (u *Undertaking4) AddDeliveryChannel() *CommunicationChannel1 {
	u.DeliveryChannel = new(CommunicationChannel1)
	return u.DeliveryChannel
}

func (u *Undertaking4) SetTransferIndicator(value string) {
	u.TransferIndicator = (*YesNoIndicator)(&value)
}

func (u *Undertaking4) AddAdditionalInformation(value string) {
	u.AdditionalInformation = append(u.AdditionalInformation, (*Max2000Text)(&value))
}
