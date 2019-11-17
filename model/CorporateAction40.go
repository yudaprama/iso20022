package model

// Provides information about the corporate action event.
type CorporateAction40 struct {

	// Provides information about the dates related to a corporate action event.
	DateDetails *CorporateActionDate58 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action event.
	PeriodDetails *CorporateActionPeriod10 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action event.
	RateAndAmountDetails *CorporateActionRate78 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action event.
	PriceDetails *CorporateActionPrice67 `xml:"PricDtls,omitempty"`

	// Provides information about securities quantity linked to a corporate action.
	SecuritiesQuantity *CorporateActionQuantity8 `xml:"SctiesQty,omitempty"`

	// Number of days used for calculating the accrued interest amount.
	InterestAccruedNumberOfDays *Max3Number `xml:"IntrstAcrdNbOfDays,omitempty"`

	// Number of the coupon attached/associated with a security.
	CouponNumber []*IdentificationFormat4Choice `xml:"CpnNb,omitempty"`

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
	DividendType *DividendTypeFormat10Choice `xml:"DvddTp,omitempty"`

	// Specifies the conversion type of an instrument.
	ConversionType *ConversionTypeFormat4Choice `xml:"ConvsTp,omitempty"`

	// Specifies the conditions in which the instructions and/or payment of the proceeds occurs.
	OccurrenceType *DistributionTypeFormat8Choice `xml:"OcrncTp,omitempty"`

	// Specifies the conditions that apply to the offer.
	OfferType []*OfferTypeFormat11Choice `xml:"OfferTp,omitempty"`

	// Specifies whether terms of the event allow resale of the rights.
	RenounceableEntitlementStatusType *RenounceableEntitlementStatusTypeFormat4Choice `xml:"RnncblEntitlmntStsTp,omitempty"`

	// Stage in the corporate action event life cycle.
	EventStage []*CorporateActionEventStageFormat20Choice `xml:"EvtStag,omitempty"`

	// Specifies the type of the additional business process linked to a corporate action event such as a claim compensation or tax refund.
	AdditionalBusinessProcessIndicator []*AdditionalBusinessProcessFormat12Choice `xml:"AddtlBizPrcInd,omitempty"`

	// Specifies the type of change announced.
	ChangeType []*CorporateActionChangeTypeFormat8Choice `xml:"ChngTp,omitempty"`

	// Type of intermediates securities distribution.
	IntermediateSecuritiesDistributionType *IntermediateSecuritiesDistributionTypeFormat18Choice `xml:"IntrmdtSctiesDstrbtnTp,omitempty"`

	// Specifies whether the capital gain is in the scope of the EU Savings directive for the income realised upon the sale, refund or redemption of shares and units (...) (Article 6(1d)).
	CapitalGainInOutIndicator *CapitalGainFormat4Choice `xml:"CptlGnInOutInd,omitempty"`

	// Specifies whether the financial instrument calculates the taxable income per dividend/taxable income per share.
	TaxableIncomePerShareCalculated *TaxableIncomePerShareCalculatedFormat4Choice `xml:"TaxblIncmPerShrClctd,omitempty"`

	// Specifies the effect on the holdings of electing a corporate action option.
	ElectionType *ElectionTypeFormat4Choice `xml:"ElctnTp,omitempty"`

	// Specifies the type of lottery announced.
	LotteryType *LotteryTypeFormat5Choice `xml:"LtryTp,omitempty"`

	// Specifies the certification format required, this is, physical or electronic format.
	CertificationType *CertificationTypeFormat4Choice `xml:"CertfctnTp,omitempty"`

	// Specifies the type of consent announced.
	ConsentType *ConsentTypeFormat5Choice `xml:"CnsntTp,omitempty"`

	// Specifies the type of information event.
	InformationType *InformationTypeFormat5Choice `xml:"InfTp,omitempty"`

	// New company's place of incorporation.
	NewPlaceOfIncorporation *RestrictedFINXMax350Text `xml:"NewPlcOfIncorprtn,omitempty"`

	// Provides additional information. This field may only be used when the information to be transmitted, cannot be coded.
	AdditionalInformation *CorporateActionNarrative39 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateAction40) AddDateDetails() *CorporateActionDate58 {
	c.DateDetails = new(CorporateActionDate58)
	return c.DateDetails
}

func (c *CorporateAction40) AddPeriodDetails() *CorporateActionPeriod10 {
	c.PeriodDetails = new(CorporateActionPeriod10)
	return c.PeriodDetails
}

func (c *CorporateAction40) AddRateAndAmountDetails() *CorporateActionRate78 {
	c.RateAndAmountDetails = new(CorporateActionRate78)
	return c.RateAndAmountDetails
}

func (c *CorporateAction40) AddPriceDetails() *CorporateActionPrice67 {
	c.PriceDetails = new(CorporateActionPrice67)
	return c.PriceDetails
}

func (c *CorporateAction40) AddSecuritiesQuantity() *CorporateActionQuantity8 {
	c.SecuritiesQuantity = new(CorporateActionQuantity8)
	return c.SecuritiesQuantity
}

func (c *CorporateAction40) SetInterestAccruedNumberOfDays(value string) {
	c.InterestAccruedNumberOfDays = (*Max3Number)(&value)
}

func (c *CorporateAction40) AddCouponNumber() *IdentificationFormat4Choice {
	newValue := new(IdentificationFormat4Choice)
	c.CouponNumber = append(c.CouponNumber, newValue)
	return newValue
}

func (c *CorporateAction40) SetCertificationBreakdownIndicator(value string) {
	c.CertificationBreakdownIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction40) SetChargesAppliedIndicator(value string) {
	c.ChargesAppliedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction40) SetRestrictionIndicator(value string) {
	c.RestrictionIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction40) SetAccruedInterestIndicator(value string) {
	c.AccruedInterestIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction40) SetLetterOfGuaranteedDeliveryIndicator(value string) {
	c.LetterOfGuaranteedDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction40) AddDividendType() *DividendTypeFormat10Choice {
	c.DividendType = new(DividendTypeFormat10Choice)
	return c.DividendType
}

func (c *CorporateAction40) AddConversionType() *ConversionTypeFormat4Choice {
	c.ConversionType = new(ConversionTypeFormat4Choice)
	return c.ConversionType
}

func (c *CorporateAction40) AddOccurrenceType() *DistributionTypeFormat8Choice {
	c.OccurrenceType = new(DistributionTypeFormat8Choice)
	return c.OccurrenceType
}

func (c *CorporateAction40) AddOfferType() *OfferTypeFormat11Choice {
	newValue := new(OfferTypeFormat11Choice)
	c.OfferType = append(c.OfferType, newValue)
	return newValue
}

func (c *CorporateAction40) AddRenounceableEntitlementStatusType() *RenounceableEntitlementStatusTypeFormat4Choice {
	c.RenounceableEntitlementStatusType = new(RenounceableEntitlementStatusTypeFormat4Choice)
	return c.RenounceableEntitlementStatusType
}

func (c *CorporateAction40) AddEventStage() *CorporateActionEventStageFormat20Choice {
	newValue := new(CorporateActionEventStageFormat20Choice)
	c.EventStage = append(c.EventStage, newValue)
	return newValue
}

func (c *CorporateAction40) AddAdditionalBusinessProcessIndicator() *AdditionalBusinessProcessFormat12Choice {
	newValue := new(AdditionalBusinessProcessFormat12Choice)
	c.AdditionalBusinessProcessIndicator = append(c.AdditionalBusinessProcessIndicator, newValue)
	return newValue
}

func (c *CorporateAction40) AddChangeType() *CorporateActionChangeTypeFormat8Choice {
	newValue := new(CorporateActionChangeTypeFormat8Choice)
	c.ChangeType = append(c.ChangeType, newValue)
	return newValue
}

func (c *CorporateAction40) AddIntermediateSecuritiesDistributionType() *IntermediateSecuritiesDistributionTypeFormat18Choice {
	c.IntermediateSecuritiesDistributionType = new(IntermediateSecuritiesDistributionTypeFormat18Choice)
	return c.IntermediateSecuritiesDistributionType
}

func (c *CorporateAction40) AddCapitalGainInOutIndicator() *CapitalGainFormat4Choice {
	c.CapitalGainInOutIndicator = new(CapitalGainFormat4Choice)
	return c.CapitalGainInOutIndicator
}

func (c *CorporateAction40) AddTaxableIncomePerShareCalculated() *TaxableIncomePerShareCalculatedFormat4Choice {
	c.TaxableIncomePerShareCalculated = new(TaxableIncomePerShareCalculatedFormat4Choice)
	return c.TaxableIncomePerShareCalculated
}

func (c *CorporateAction40) AddElectionType() *ElectionTypeFormat4Choice {
	c.ElectionType = new(ElectionTypeFormat4Choice)
	return c.ElectionType
}

func (c *CorporateAction40) AddLotteryType() *LotteryTypeFormat5Choice {
	c.LotteryType = new(LotteryTypeFormat5Choice)
	return c.LotteryType
}

func (c *CorporateAction40) AddCertificationType() *CertificationTypeFormat4Choice {
	c.CertificationType = new(CertificationTypeFormat4Choice)
	return c.CertificationType
}

func (c *CorporateAction40) AddConsentType() *ConsentTypeFormat5Choice {
	c.ConsentType = new(ConsentTypeFormat5Choice)
	return c.ConsentType
}

func (c *CorporateAction40) AddInformationType() *InformationTypeFormat5Choice {
	c.InformationType = new(InformationTypeFormat5Choice)
	return c.InformationType
}

func (c *CorporateAction40) SetNewPlaceOfIncorporation(value string) {
	c.NewPlaceOfIncorporation = (*RestrictedFINXMax350Text)(&value)
}

func (c *CorporateAction40) AddAdditionalInformation() *CorporateActionNarrative39 {
	c.AdditionalInformation = new(CorporateActionNarrative39)
	return c.AdditionalInformation
}
