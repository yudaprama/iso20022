package model

// Provides information about the corporate action option.
type CorporateActionOption115 struct {

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption18Choice `xml:"OptnTp"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType26Choice `xml:"FrctnDspstn,omitempty"`

	// Specifies the conditions that apply to the offer.
	OfferType []*OfferTypeFormat10Choice `xml:"OfferTp,omitempty"`

	// Specifies the features that may apply to a corporate action option.
	OptionFeatures []*OptionFeaturesFormat17Choice `xml:"OptnFeatrs,omitempty"`

	// Specifies the type of intermediates securities distribution.
	IntermediateSecuritiesDistributionType *IntermediateSecuritiesDistributionTypeFormat15Choice `xml:"IntrmdtSctiesDstrbtnTp,omitempty"`

	// Specifies the status of the option.
	OptionAvailabilityStatus *OptionAvailabilityStatus3Choice `xml:"OptnAvlbtySts,omitempty"`

	// Indicates the type of certification/breakdown.
	CertificationBreakdownType []*BeneficiaryCertificationType9Choice `xml:"CertfctnBrkdwnTp,omitempty"`

	// Holder of the security has to certify, in line with the terms of the corporate action, that it is not domiciled in the country indicated.
	NonDomicileCountry []*CountryCode `xml:"NonDmclCtry,omitempty"`

	// Country of domicile in which the Corporate Action option is valid. The holder of the security has to certify that it is domiciled in the country indicated.
	ValidDomicileCountry []*CountryCode `xml:"VldDmclCtry,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Indicates whether the corporate action movement is a default processing or a standing instruction.
	DefaultProcessingOrStandingInstruction *DefaultProcessingOrStandingInstruction1Choice `xml:"DfltPrcgOrStgInstr"`

	// Indicates whether redemption charges apply.
	ChargesAppliedIndicator *YesNoIndicator `xml:"ChrgsApldInd,omitempty"`

	// Indicates whether or not certification/breakdown is required from the account owner.
	// Yes: certification required
	// No: no certification required
	CertificationBreakdownIndicator *YesNoIndicator `xml:"CertfctnBrkdwnInd,omitempty"`

	// Indicates whether withdrawal of instruction is allowed.
	WithdrawalAllowedIndicator *YesNoIndicator `xml:"WdrwlAllwdInd,omitempty"`

	// Indicates whether change of instruction is allowed.
	ChangeAllowedIndicator *YesNoIndicator `xml:"ChngAllwdInd,omitempty"`

	// Indicates whether the option, different from the default one, shall be applied by the account owner.
	AppliedOptionIndicator *YesNoIndicator `xml:"ApldOptnInd,omitempty"`

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification19 `xml:"SctyId,omitempty"`

	// Provides information about the dates related to a corporate action option.
	DateDetails *CorporateActionDate48 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action option.
	PeriodDetails *CorporateActionPeriod7 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate67 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice58 `xml:"PricDtls,omitempty"`

	// Provides information about securities quantity related to a corporate action option.
	SecuritiesQuantity *SecuritiesOption51 `xml:"SctiesQty,omitempty"`

	// Provides information about securities movement related to a corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption49 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement linked to the corporate action option.
	CashMovementDetails []*CashOption42 `xml:"CshMvmntDtls,omitempty"`

	// Provides additional information about the corporate action movement.
	AdditionalInformation *CorporateActionNarrative29 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionOption115) SetOptionNumber(value string) {
	c.OptionNumber = (*Exact3NumericText)(&value)
}

func (c *CorporateActionOption115) AddOptionType() *CorporateActionOption18Choice {
	c.OptionType = new(CorporateActionOption18Choice)
	return c.OptionType
}

func (c *CorporateActionOption115) AddFractionDisposition() *FractionDispositionType26Choice {
	c.FractionDisposition = new(FractionDispositionType26Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption115) AddOfferType() *OfferTypeFormat10Choice {
	newValue := new(OfferTypeFormat10Choice)
	c.OfferType = append(c.OfferType, newValue)
	return newValue
}

func (c *CorporateActionOption115) AddOptionFeatures() *OptionFeaturesFormat17Choice {
	newValue := new(OptionFeaturesFormat17Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption115) AddIntermediateSecuritiesDistributionType() *IntermediateSecuritiesDistributionTypeFormat15Choice {
	c.IntermediateSecuritiesDistributionType = new(IntermediateSecuritiesDistributionTypeFormat15Choice)
	return c.IntermediateSecuritiesDistributionType
}

func (c *CorporateActionOption115) AddOptionAvailabilityStatus() *OptionAvailabilityStatus3Choice {
	c.OptionAvailabilityStatus = new(OptionAvailabilityStatus3Choice)
	return c.OptionAvailabilityStatus
}

func (c *CorporateActionOption115) AddCertificationBreakdownType() *BeneficiaryCertificationType9Choice {
	newValue := new(BeneficiaryCertificationType9Choice)
	c.CertificationBreakdownType = append(c.CertificationBreakdownType, newValue)
	return newValue
}

func (c *CorporateActionOption115) AddNonDomicileCountry(value string) {
	c.NonDomicileCountry = append(c.NonDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption115) AddValidDomicileCountry(value string) {
	c.ValidDomicileCountry = append(c.ValidDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption115) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption115) AddDefaultProcessingOrStandingInstruction() *DefaultProcessingOrStandingInstruction1Choice {
	c.DefaultProcessingOrStandingInstruction = new(DefaultProcessingOrStandingInstruction1Choice)
	return c.DefaultProcessingOrStandingInstruction
}

func (c *CorporateActionOption115) SetChargesAppliedIndicator(value string) {
	c.ChargesAppliedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption115) SetCertificationBreakdownIndicator(value string) {
	c.CertificationBreakdownIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption115) SetWithdrawalAllowedIndicator(value string) {
	c.WithdrawalAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption115) SetChangeAllowedIndicator(value string) {
	c.ChangeAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption115) SetAppliedOptionIndicator(value string) {
	c.AppliedOptionIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption115) AddSecurityIdentification() *SecurityIdentification19 {
	c.SecurityIdentification = new(SecurityIdentification19)
	return c.SecurityIdentification
}

func (c *CorporateActionOption115) AddDateDetails() *CorporateActionDate48 {
	c.DateDetails = new(CorporateActionDate48)
	return c.DateDetails
}

func (c *CorporateActionOption115) AddPeriodDetails() *CorporateActionPeriod7 {
	c.PeriodDetails = new(CorporateActionPeriod7)
	return c.PeriodDetails
}

func (c *CorporateActionOption115) AddRateAndAmountDetails() *CorporateActionRate67 {
	c.RateAndAmountDetails = new(CorporateActionRate67)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption115) AddPriceDetails() *CorporateActionPrice58 {
	c.PriceDetails = new(CorporateActionPrice58)
	return c.PriceDetails
}

func (c *CorporateActionOption115) AddSecuritiesQuantity() *SecuritiesOption51 {
	c.SecuritiesQuantity = new(SecuritiesOption51)
	return c.SecuritiesQuantity
}

func (c *CorporateActionOption115) AddSecuritiesMovementDetails() *SecuritiesOption49 {
	newValue := new(SecuritiesOption49)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption115) AddCashMovementDetails() *CashOption42 {
	newValue := new(CashOption42)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption115) AddAdditionalInformation() *CorporateActionNarrative29 {
	c.AdditionalInformation = new(CorporateActionNarrative29)
	return c.AdditionalInformation
}
