package model

// Provides information about the CA option.
type CorporateActionOption1 struct {

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption1FormatChoice `xml:"OptnTp"`

	// Specifies the status of the option.
	OptionAvailabilityStatus *CorporateActionEventStatus2FormatChoice `xml:"OptnAvlbtySts"`

	// Whether or not certification is required from the account owner.
	// Yes: certification required
	// No: no certification required
	CertificationIndicator *YesNoIndicator `xml:"CertfctnInd,omitempty"`

	// Type of certification which is required.
	CertificationType *BeneficiaryCertificationType1FormatChoice `xml:"CertfctnTp,omitempty"`

	// Identification of a temporary security used for processing reasons, eg, contra security used in the US.
	AssentedLineSecurityIdentification *SecurityIdentification7 `xml:"AssntdLineSctyId,omitempty"`

	// Identification of the safekeeping account held by an agent at the CSD.
	AgentSecuritiesAccountIdentification *Max35Text `xml:"AgtSctiesAcctId,omitempty"`

	// Identification of the cash account held by an agent at the CSD.
	AgentCashAccountIdentification *AccountIdentification2Choice `xml:"AgtCshAcctId,omitempty"`

	// Specifies the conditions that apply to the offer.
	OfferType []*OfferType1FormatChoice `xml:"OfferTp,omitempty"`

	// Type of intermediates securities distribution, eg, stock dividend, reverse right.
	IntermediateSecuritiesDistributionType *IntermediateSecurityDistributionType1FormatChoice `xml:"IntrmdtSctiesDstrbtnTp,omitempty"`

	// Indicates whether withdrawal of instruction is allowed.
	WithdrawalAllowedIndicator *YesNoIndicator `xml:"WdrwlAllwdInd"`

	// Indicates whether change of instruction is allowed.
	ChangeAllowedIndicator *YesNoIndicator `xml:"ChngAllwdInd"`

	// Provides information about the dates related to a CA option.
	DateDetails *CorporateActionDate4 `xml:"DtDtls,omitempty"`

	// Provides information about rates and amounts related to a CA option.
	RateAndAmountDetails *CorporateActionRate2 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a CA option.
	PriceDetails *CorporateActionPrice1 `xml:"PricDtls,omitempty"`

	// Provides information about the periods related to a CA option.
	PeriodDetails *CorporateActionPeriod2 `xml:"PrdDtls,omitempty"`

	// Provides information about the securities movement linked to the CA option.
	SecuritiesMovementDetails []*SecurityOption1 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement linked to the CA option.
	CashMovementDetails []*CashOption1 `xml:"CshMvmntDtls,omitempty"`

	// Provides information about the agents linked to the CA option.
	CorporateActionOtherAgentDetails []*CorporateActionAgent1 `xml:"CorpActnOthrAgtDtls,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType1FormatChoice `xml:"FrctnDspstn,omitempty"`

	// ndicates whether redemption charges apply.
	RedemptionChargesAppliedIndicator *YesNoIndicator `xml:"RedChrgsApldInd,omitempty"`

	// Specifies the features that may apply to a corporate action option.
	OptionFeatures []*OptionFeatures1FormatChoice `xml:"OptnFeatrs,omitempty"`

	// Provides additional information.
	CorporateActionAdditionalInformation *CorporateActionNarrative1 `xml:"CorpActnAddtlInf,omitempty"`
}

func (c *CorporateActionOption1) SetOptionNumber(value string) {
	c.OptionNumber = (*Exact3NumericText)(&value)
}

func (c *CorporateActionOption1) AddOptionType() *CorporateActionOption1FormatChoice {
	c.OptionType = new(CorporateActionOption1FormatChoice)
	return c.OptionType
}

func (c *CorporateActionOption1) AddOptionAvailabilityStatus() *CorporateActionEventStatus2FormatChoice {
	c.OptionAvailabilityStatus = new(CorporateActionEventStatus2FormatChoice)
	return c.OptionAvailabilityStatus
}

func (c *CorporateActionOption1) SetCertificationIndicator(value string) {
	c.CertificationIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption1) AddCertificationType() *BeneficiaryCertificationType1FormatChoice {
	c.CertificationType = new(BeneficiaryCertificationType1FormatChoice)
	return c.CertificationType
}

func (c *CorporateActionOption1) AddAssentedLineSecurityIdentification() *SecurityIdentification7 {
	c.AssentedLineSecurityIdentification = new(SecurityIdentification7)
	return c.AssentedLineSecurityIdentification
}

func (c *CorporateActionOption1) SetAgentSecuritiesAccountIdentification(value string) {
	c.AgentSecuritiesAccountIdentification = (*Max35Text)(&value)
}

func (c *CorporateActionOption1) AddAgentCashAccountIdentification() *AccountIdentification2Choice {
	c.AgentCashAccountIdentification = new(AccountIdentification2Choice)
	return c.AgentCashAccountIdentification
}

func (c *CorporateActionOption1) AddOfferType() *OfferType1FormatChoice {
	newValue := new(OfferType1FormatChoice)
	c.OfferType = append(c.OfferType, newValue)
	return newValue
}

func (c *CorporateActionOption1) AddIntermediateSecuritiesDistributionType() *IntermediateSecurityDistributionType1FormatChoice {
	c.IntermediateSecuritiesDistributionType = new(IntermediateSecurityDistributionType1FormatChoice)
	return c.IntermediateSecuritiesDistributionType
}

func (c *CorporateActionOption1) SetWithdrawalAllowedIndicator(value string) {
	c.WithdrawalAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption1) SetChangeAllowedIndicator(value string) {
	c.ChangeAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption1) AddDateDetails() *CorporateActionDate4 {
	c.DateDetails = new(CorporateActionDate4)
	return c.DateDetails
}

func (c *CorporateActionOption1) AddRateAndAmountDetails() *CorporateActionRate2 {
	c.RateAndAmountDetails = new(CorporateActionRate2)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption1) AddPriceDetails() *CorporateActionPrice1 {
	c.PriceDetails = new(CorporateActionPrice1)
	return c.PriceDetails
}

func (c *CorporateActionOption1) AddPeriodDetails() *CorporateActionPeriod2 {
	c.PeriodDetails = new(CorporateActionPeriod2)
	return c.PeriodDetails
}

func (c *CorporateActionOption1) AddSecuritiesMovementDetails() *SecurityOption1 {
	newValue := new(SecurityOption1)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption1) AddCashMovementDetails() *CashOption1 {
	newValue := new(CashOption1)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption1) AddCorporateActionOtherAgentDetails() *CorporateActionAgent1 {
	newValue := new(CorporateActionAgent1)
	c.CorporateActionOtherAgentDetails = append(c.CorporateActionOtherAgentDetails, newValue)
	return newValue
}

func (c *CorporateActionOption1) AddFractionDisposition() *FractionDispositionType1FormatChoice {
	c.FractionDisposition = new(FractionDispositionType1FormatChoice)
	return c.FractionDisposition
}

func (c *CorporateActionOption1) SetRedemptionChargesAppliedIndicator(value string) {
	c.RedemptionChargesAppliedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption1) AddOptionFeatures() *OptionFeatures1FormatChoice {
	newValue := new(OptionFeatures1FormatChoice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption1) AddCorporateActionAdditionalInformation() *CorporateActionNarrative1 {
	c.CorporateActionAdditionalInformation = new(CorporateActionNarrative1)
	return c.CorporateActionAdditionalInformation
}
