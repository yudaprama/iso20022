package model

// Independent undertaking, such as a demand guarantee or standby letter of credit, that provides financial assurance, to be honoured on the presentation of documents that comply with its terms and conditions.
type Undertaking2 struct {

	// Undertaking name.
	Name *UndertakingName1Code `xml:"Nm,omitempty"`

	// Party in whose favour the counter-undertaking is issued.
	Beneficiary *PartyIdentification43 `xml:"Bnfcry,omitempty"`

	// Details related to the expiry terms of the counter-undertaking.
	ExpiryDetails *ExpiryDetails2 `xml:"XpryDtls,omitempty"`

	// Details related to the amount of the counter-undertaking.
	CounterUndertakingAmount *UndertakingAmount1 `xml:"CntrUdrtkgAmt,omitempty"`

	// Indicates whether the applicant/obligor or beneficiary is responsible for payment of the confirmation charges.
	ConfirmationChargesPayableBy *ExternalTypeOfParty1Code `xml:"ConfChrgsPyblBy,omitempty"`

	// Rules and laws governing the counter-undertaking.
	GovernanceRulesAndLaw *GovernanceRules1 `xml:"GovncRulesAndLaw,omitempty"`

	// Indication as to whether a claim is to utilise a standard claim form of the issuing institution.
	StandardClaimDocumentIndicator *YesNoIndicator `xml:"StdClmDocInd,omitempty"`

	// Additional information related to the counter-undertaking.
	AdditionalInformation []*Max2000Text `xml:"AddtlInf,omitempty"`
}

func (u *Undertaking2) SetName(value string) {
	u.Name = (*UndertakingName1Code)(&value)
}

func (u *Undertaking2) AddBeneficiary() *PartyIdentification43 {
	u.Beneficiary = new(PartyIdentification43)
	return u.Beneficiary
}

func (u *Undertaking2) AddExpiryDetails() *ExpiryDetails2 {
	u.ExpiryDetails = new(ExpiryDetails2)
	return u.ExpiryDetails
}

func (u *Undertaking2) AddCounterUndertakingAmount() *UndertakingAmount1 {
	u.CounterUndertakingAmount = new(UndertakingAmount1)
	return u.CounterUndertakingAmount
}

func (u *Undertaking2) SetConfirmationChargesPayableBy(value string) {
	u.ConfirmationChargesPayableBy = (*ExternalTypeOfParty1Code)(&value)
}

func (u *Undertaking2) AddGovernanceRulesAndLaw() *GovernanceRules1 {
	u.GovernanceRulesAndLaw = new(GovernanceRules1)
	return u.GovernanceRulesAndLaw
}

func (u *Undertaking2) SetStandardClaimDocumentIndicator(value string) {
	u.StandardClaimDocumentIndicator = (*YesNoIndicator)(&value)
}

func (u *Undertaking2) AddAdditionalInformation(value string) {
	u.AdditionalInformation = append(u.AdditionalInformation, (*Max2000Text)(&value))
}
