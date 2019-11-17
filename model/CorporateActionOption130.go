package model

// Provides information about the corporate action option.
type CorporateActionOption130 struct {

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption18Choice `xml:"OptnTp"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType26Choice `xml:"FrctnDspstn,omitempty"`

	// Specifies the conditions that apply to the offer.
	OfferType []*OfferTypeFormat10Choice `xml:"OfferTp,omitempty"`

	// Specifies the features that may apply to a corporate action option.
	OptionFeatures []*OptionFeaturesFormat22Choice `xml:"OptnFeatrs,omitempty"`

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

	// Indicates whether charges apply to the holder, for instance redemption charges.
	ChargesAppliedIndicator *YesNoIndicator `xml:"ChrgsApldInd,omitempty"`

	// Whether or not certification/breakdown is required from the account owner.
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
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId,omitempty"`

	// Provides information about the dates related to a corporate action option.
	DateDetails *CorporateActionDate48 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action option.
	PeriodDetails *CorporateActionPeriod7 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate80 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice58 `xml:"PricDtls,omitempty"`

	// Provides information about securities quantity linked to a corporate action option.
	SecuritiesQuantity *SecuritiesOption51 `xml:"SctiesQty,omitempty"`

	// Provides information about the securities movement linked to the corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption59 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement linked to the corporate action option.
	CashMovementDetails []*CashOption50 `xml:"CshMvmntDtls,omitempty"`

	// Provides additional information.
	AdditionalInformation *CorporateActionNarrative29 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionOption130) SetOptionNumber(value string) {
	c.OptionNumber = (*Exact3NumericText)(&value)
}

func (c *CorporateActionOption130) AddOptionType() *CorporateActionOption18Choice {
	c.OptionType = new(CorporateActionOption18Choice)
	return c.OptionType
}

func (c *CorporateActionOption130) AddFractionDisposition() *FractionDispositionType26Choice {
	c.FractionDisposition = new(FractionDispositionType26Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption130) AddOfferType() *OfferTypeFormat10Choice {
	newValue := new(OfferTypeFormat10Choice)
	c.OfferType = append(c.OfferType, newValue)
	return newValue
}

func (c *CorporateActionOption130) AddOptionFeatures() *OptionFeaturesFormat22Choice {
	newValue := new(OptionFeaturesFormat22Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption130) AddOptionAvailabilityStatus() *OptionAvailabilityStatus3Choice {
	c.OptionAvailabilityStatus = new(OptionAvailabilityStatus3Choice)
	return c.OptionAvailabilityStatus
}

func (c *CorporateActionOption130) AddCertificationBreakdownType() *BeneficiaryCertificationType9Choice {
	newValue := new(BeneficiaryCertificationType9Choice)
	c.CertificationBreakdownType = append(c.CertificationBreakdownType, newValue)
	return newValue
}

func (c *CorporateActionOption130) AddNonDomicileCountry(value string) {
	c.NonDomicileCountry = append(c.NonDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption130) AddValidDomicileCountry(value string) {
	c.ValidDomicileCountry = append(c.ValidDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption130) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption130) AddDefaultProcessingOrStandingInstruction() *DefaultProcessingOrStandingInstruction1Choice {
	c.DefaultProcessingOrStandingInstruction = new(DefaultProcessingOrStandingInstruction1Choice)
	return c.DefaultProcessingOrStandingInstruction
}

func (c *CorporateActionOption130) SetChargesAppliedIndicator(value string) {
	c.ChargesAppliedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption130) SetCertificationBreakdownIndicator(value string) {
	c.CertificationBreakdownIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption130) SetWithdrawalAllowedIndicator(value string) {
	c.WithdrawalAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption130) SetChangeAllowedIndicator(value string) {
	c.ChangeAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption130) SetAppliedOptionIndicator(value string) {
	c.AppliedOptionIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption130) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	c.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return c.FinancialInstrumentIdentification
}

func (c *CorporateActionOption130) AddDateDetails() *CorporateActionDate48 {
	c.DateDetails = new(CorporateActionDate48)
	return c.DateDetails
}

func (c *CorporateActionOption130) AddPeriodDetails() *CorporateActionPeriod7 {
	c.PeriodDetails = new(CorporateActionPeriod7)
	return c.PeriodDetails
}

func (c *CorporateActionOption130) AddRateAndAmountDetails() *CorporateActionRate80 {
	c.RateAndAmountDetails = new(CorporateActionRate80)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption130) AddPriceDetails() *CorporateActionPrice58 {
	c.PriceDetails = new(CorporateActionPrice58)
	return c.PriceDetails
}

func (c *CorporateActionOption130) AddSecuritiesQuantity() *SecuritiesOption51 {
	c.SecuritiesQuantity = new(SecuritiesOption51)
	return c.SecuritiesQuantity
}

func (c *CorporateActionOption130) AddSecuritiesMovementDetails() *SecuritiesOption59 {
	newValue := new(SecuritiesOption59)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption130) AddCashMovementDetails() *CashOption50 {
	newValue := new(CashOption50)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption130) AddAdditionalInformation() *CorporateActionNarrative29 {
	c.AdditionalInformation = new(CorporateActionNarrative29)
	return c.AdditionalInformation
}
