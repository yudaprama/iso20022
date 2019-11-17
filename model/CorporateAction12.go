package model

// Provides information about the corporate action event.
type CorporateAction12 struct {

	// Provides information about the dates related to a corporate action event.
	DateDetails *CorporateActionDate27 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action event.
	PeriodDetails *CorporateActionPeriod10 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action event.
	RateAndAmountDetails *CorporateActionRate43 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action event.
	PriceDetails *CorporateActionPrice42 `xml:"PricDtls,omitempty"`

	// Provides information about securities quantity linked to a corporate action.
	SecuritiesQuantity *CorporateActionQuantity5 `xml:"SctiesQty,omitempty"`

	// Number of days used for calculating the accrued interest amount.
	InterestAccruedNumberOfDays *Max3Number `xml:"IntrstAcrdNbOfDays,omitempty"`

	// Number of the coupon attached/associated with a security.
	CouponNumber []*IdentificationFormat1Choice `xml:"CpnNb,omitempty"`

	// Indicates whether certification/breakdown is required.
	// Yes = certification required.
	// No = no certification required.
	CertificationBreakdownIndicator *YesNoIndicator `xml:"CertfctnBrkdwnInd,omitempty"`

	// Indicates whether charges apply to the holder, for instance redemption charges.
	ChargesAppliedIndicator *YesNoIndicator `xml:"ChrgsApldInd,omitempty"`

	// Indicates whether there is restrictions apply to the corporate action event or not.
	// Yes = There is restrictions.
	// No = There is no restrictions.
	RestrictionIndicator *YesNoIndicator `xml:"RstrctnInd,omitempty"`

	// Indicates whether the holder is entitled to accrued interest.
	AccruedInterestIndicator *YesNoIndicator `xml:"AcrdIntrstInd,omitempty"`

	// Indicates whether a letter of guaranteed delivery can be submitted in order to participate in the offer on full eligible position. It is not intended for use in situations arising from failed or late trades.
	LetterOfGuaranteedDeliveryIndicator *YesNoIndicator `xml:"LttrOfGrntedDlvryInd,omitempty"`

	// Specifies the conditions in which a dividend is paid.
	DividendType *DividendTypeFormat5Choice `xml:"DvddTp,omitempty"`

	// Specifies the conversion type of an instrument.
	ConversionType *ConversionTypeFormat1Choice `xml:"ConvsTp,omitempty"`

	// Specifies the conditions in which the instructions and/or payment of the proceeds occurs.
	OccurrenceType *DistributionTypeFormat4Choice `xml:"OcrncTp,omitempty"`

	// Specifies the conditions that apply to the offer.
	OfferType []*OfferTypeFormat5Choice `xml:"OfferTp,omitempty"`

	// Specifies whether terms of the event allow resale of the rights.
	RenounceableEntitlementStatusType *RenounceableEntitlementStatusTypeFormat1Choice `xml:"RnncblEntitlmntStsTp,omitempty"`

	// Stage in the corporate action event life cycle.
	EventStage []*CorporateActionEventStageFormat5Choice `xml:"EvtStag,omitempty"`

	// Specifies the type of the additional business process linked to a corporate action event such as a claim compensation or tax refund.
	AdditionalBusinessProcessIndicator []*AdditionalBusinessProcessFormat1Choice `xml:"AddtlBizPrcInd,omitempty"`

	// Specifies the type of change announced.
	ChangeType []*CorporateActionChangeTypeFormat1Choice `xml:"ChngTp,omitempty"`

	// Type of intermediates securities distribution.
	IntermediateSecuritiesDistributionType *IntermediateSecuritiesDistributionTypeFormat9Choice `xml:"IntrmdtSctiesDstrbtnTp,omitempty"`

	// Specifies whether the capital gain is in the scope of the EU Savings directive for the income realised upon the sale, refund or redemption of shares and units (...) (Article 6(1d)).
	CapitalGainInOutIndicator *CapitalGainFormat1Choice `xml:"CptlGnInOutInd,omitempty"`

	// Specifies whether the financial instrument calculates the taxable income per dividend/taxable income per share.
	TaxableIncomePerShareCalculated *TaxableIncomePerShareCalculatedFormat1Choice `xml:"TaxblIncmPerShrClctd,omitempty"`

	// Specifies the effect on the holdings of electing a corporate action option.
	ElectionType *ElectionTypeFormat1Choice `xml:"ElctnTp,omitempty"`

	// Specifies the type of lottery announced.
	LotteryType *LotteryTypeFormat1Choice `xml:"LtryTp,omitempty"`

	// Specifies the certification format required, this is, physical or electronic format.
	CertificationType *CertificationTypeFormat1Choice `xml:"CertfctnTp,omitempty"`

	// Specifies the type of consent announced.
	ConsentType *ConsentTypeFormat1Choice `xml:"CnsntTp,omitempty"`

	// Specifies the type of information event.
	InformationType *InformationTypeFormat1Choice `xml:"InfTp,omitempty"`

	// New company's place of incorporation.
	NewPlaceOfIncorporation *Max350Text `xml:"NewPlcOfIncorprtn,omitempty"`

	// Provides additional information. This field may only be used when the information to be transmitted, cannot be coded.
	AdditionalInformation *CorporateActionNarrative24 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateAction12) AddDateDetails() *CorporateActionDate27 {
	c.DateDetails = new(CorporateActionDate27)
	return c.DateDetails
}

func (c *CorporateAction12) AddPeriodDetails() *CorporateActionPeriod10 {
	c.PeriodDetails = new(CorporateActionPeriod10)
	return c.PeriodDetails
}

func (c *CorporateAction12) AddRateAndAmountDetails() *CorporateActionRate43 {
	c.RateAndAmountDetails = new(CorporateActionRate43)
	return c.RateAndAmountDetails
}

func (c *CorporateAction12) AddPriceDetails() *CorporateActionPrice42 {
	c.PriceDetails = new(CorporateActionPrice42)
	return c.PriceDetails
}

func (c *CorporateAction12) AddSecuritiesQuantity() *CorporateActionQuantity5 {
	c.SecuritiesQuantity = new(CorporateActionQuantity5)
	return c.SecuritiesQuantity
}

func (c *CorporateAction12) SetInterestAccruedNumberOfDays(value string) {
	c.InterestAccruedNumberOfDays = (*Max3Number)(&value)
}

func (c *CorporateAction12) AddCouponNumber() *IdentificationFormat1Choice {
	newValue := new(IdentificationFormat1Choice)
	c.CouponNumber = append(c.CouponNumber, newValue)
	return newValue
}

func (c *CorporateAction12) SetCertificationBreakdownIndicator(value string) {
	c.CertificationBreakdownIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction12) SetChargesAppliedIndicator(value string) {
	c.ChargesAppliedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction12) SetRestrictionIndicator(value string) {
	c.RestrictionIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction12) SetAccruedInterestIndicator(value string) {
	c.AccruedInterestIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction12) SetLetterOfGuaranteedDeliveryIndicator(value string) {
	c.LetterOfGuaranteedDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction12) AddDividendType() *DividendTypeFormat5Choice {
	c.DividendType = new(DividendTypeFormat5Choice)
	return c.DividendType
}

func (c *CorporateAction12) AddConversionType() *ConversionTypeFormat1Choice {
	c.ConversionType = new(ConversionTypeFormat1Choice)
	return c.ConversionType
}

func (c *CorporateAction12) AddOccurrenceType() *DistributionTypeFormat4Choice {
	c.OccurrenceType = new(DistributionTypeFormat4Choice)
	return c.OccurrenceType
}

func (c *CorporateAction12) AddOfferType() *OfferTypeFormat5Choice {
	newValue := new(OfferTypeFormat5Choice)
	c.OfferType = append(c.OfferType, newValue)
	return newValue
}

func (c *CorporateAction12) AddRenounceableEntitlementStatusType() *RenounceableEntitlementStatusTypeFormat1Choice {
	c.RenounceableEntitlementStatusType = new(RenounceableEntitlementStatusTypeFormat1Choice)
	return c.RenounceableEntitlementStatusType
}

func (c *CorporateAction12) AddEventStage() *CorporateActionEventStageFormat5Choice {
	newValue := new(CorporateActionEventStageFormat5Choice)
	c.EventStage = append(c.EventStage, newValue)
	return newValue
}

func (c *CorporateAction12) AddAdditionalBusinessProcessIndicator() *AdditionalBusinessProcessFormat1Choice {
	newValue := new(AdditionalBusinessProcessFormat1Choice)
	c.AdditionalBusinessProcessIndicator = append(c.AdditionalBusinessProcessIndicator, newValue)
	return newValue
}

func (c *CorporateAction12) AddChangeType() *CorporateActionChangeTypeFormat1Choice {
	newValue := new(CorporateActionChangeTypeFormat1Choice)
	c.ChangeType = append(c.ChangeType, newValue)
	return newValue
}

func (c *CorporateAction12) AddIntermediateSecuritiesDistributionType() *IntermediateSecuritiesDistributionTypeFormat9Choice {
	c.IntermediateSecuritiesDistributionType = new(IntermediateSecuritiesDistributionTypeFormat9Choice)
	return c.IntermediateSecuritiesDistributionType
}

func (c *CorporateAction12) AddCapitalGainInOutIndicator() *CapitalGainFormat1Choice {
	c.CapitalGainInOutIndicator = new(CapitalGainFormat1Choice)
	return c.CapitalGainInOutIndicator
}

func (c *CorporateAction12) AddTaxableIncomePerShareCalculated() *TaxableIncomePerShareCalculatedFormat1Choice {
	c.TaxableIncomePerShareCalculated = new(TaxableIncomePerShareCalculatedFormat1Choice)
	return c.TaxableIncomePerShareCalculated
}

func (c *CorporateAction12) AddElectionType() *ElectionTypeFormat1Choice {
	c.ElectionType = new(ElectionTypeFormat1Choice)
	return c.ElectionType
}

func (c *CorporateAction12) AddLotteryType() *LotteryTypeFormat1Choice {
	c.LotteryType = new(LotteryTypeFormat1Choice)
	return c.LotteryType
}

func (c *CorporateAction12) AddCertificationType() *CertificationTypeFormat1Choice {
	c.CertificationType = new(CertificationTypeFormat1Choice)
	return c.CertificationType
}

func (c *CorporateAction12) AddConsentType() *ConsentTypeFormat1Choice {
	c.ConsentType = new(ConsentTypeFormat1Choice)
	return c.ConsentType
}

func (c *CorporateAction12) AddInformationType() *InformationTypeFormat1Choice {
	c.InformationType = new(InformationTypeFormat1Choice)
	return c.InformationType
}

func (c *CorporateAction12) SetNewPlaceOfIncorporation(value string) {
	c.NewPlaceOfIncorporation = (*Max350Text)(&value)
}

func (c *CorporateAction12) AddAdditionalInformation() *CorporateActionNarrative24 {
	c.AdditionalInformation = new(CorporateActionNarrative24)
	return c.AdditionalInformation
}
