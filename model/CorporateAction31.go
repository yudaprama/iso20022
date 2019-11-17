package model

// Provides information about the corporate action event.
type CorporateAction31 struct {

	// Provides information about the dates related to a corporate action event.
	DateDetails *CorporateActionDate44 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action event.
	PeriodDetails *CorporateActionPeriod10 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action event.
	RateAndAmountDetails *CorporateActionRate66 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action event.
	PriceDetails *CorporateActionPrice57 `xml:"PricDtls,omitempty"`

	// Provides information about securities quantity linked to a corporate action.
	SecuritiesQuantity *CorporateActionQuantity7 `xml:"SctiesQty,omitempty"`

	// Number of days used for calculating the accrued interest amount.
	InterestAccruedNumberOfDays *Max3Number `xml:"IntrstAcrdNbOfDays,omitempty"`

	// Number of the coupon attached/associated with a security.
	CouponNumber []*IdentificationFormat3Choice `xml:"CpnNb,omitempty"`

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
	DividendType *DividendTypeFormat9Choice `xml:"DvddTp,omitempty"`

	// Specifies the conversion type of an instrument.
	ConversionType *ConversionTypeFormat3Choice `xml:"ConvsTp,omitempty"`

	// Specifies the conditions in which the instructions and/or payment of the proceeds occurs.
	OccurrenceType *DistributionTypeFormat7Choice `xml:"OcrncTp,omitempty"`

	// Specifies the conditions that apply to the offer.
	OfferType []*OfferTypeFormat10Choice `xml:"OfferTp,omitempty"`

	// Specifies whether terms of the event allow resale of the rights.
	RenounceableEntitlementStatusType *RenounceableEntitlementStatusTypeFormat3Choice `xml:"RnncblEntitlmntStsTp,omitempty"`

	// Stage in the corporate action event life cycle.
	EventStage []*CorporateActionEventStageFormat13Choice `xml:"EvtStag,omitempty"`

	// Specifies the type of the additional business process linked to a corporate action event such as a claim compensation or tax refund.
	AdditionalBusinessProcessIndicator []*AdditionalBusinessProcessFormat9Choice `xml:"AddtlBizPrcInd,omitempty"`

	// Specifies the type of change announced.
	ChangeType []*CorporateActionChangeTypeFormat5Choice `xml:"ChngTp,omitempty"`

	// Type of intermediates securities distribution.
	IntermediateSecuritiesDistributionType *IntermediateSecuritiesDistributionTypeFormat15Choice `xml:"IntrmdtSctiesDstrbtnTp,omitempty"`

	// Specifies whether the capital gain is in the scope of the EU Savings directive for the income realised upon the sale, refund or redemption of shares and units (...) (Article 6(1d)).
	CapitalGainInOutIndicator *CapitalGainFormat3Choice `xml:"CptlGnInOutInd,omitempty"`

	// Specifies whether the financial instrument calculates the taxable income per dividend/taxable income per share.
	TaxableIncomePerShareCalculated *TaxableIncomePerShareCalculatedFormat3Choice `xml:"TaxblIncmPerShrClctd,omitempty"`

	// Specifies the effect on the holdings of electing a corporate action option.
	ElectionType *ElectionTypeFormat3Choice `xml:"ElctnTp,omitempty"`

	// Specifies the type of lottery announced.
	LotteryType *LotteryTypeFormat4Choice `xml:"LtryTp,omitempty"`

	// Specifies the certification format required, this is, physical or electronic format.
	CertificationType *CertificationTypeFormat3Choice `xml:"CertfctnTp,omitempty"`

	// Specifies the type of consent announced.
	ConsentType *ConsentTypeFormat4Choice `xml:"CnsntTp,omitempty"`

	// Specifies the type of information event.
	InformationType *InformationTypeFormat4Choice `xml:"InfTp,omitempty"`

	// New company's place of incorporation.
	NewPlaceOfIncorporation *Max350Text `xml:"NewPlcOfIncorprtn,omitempty"`

	// Provides additional information. This field may only be used when the information to be transmitted, cannot be coded.
	AdditionalInformation *CorporateActionNarrative26 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateAction31) AddDateDetails() *CorporateActionDate44 {
	c.DateDetails = new(CorporateActionDate44)
	return c.DateDetails
}

func (c *CorporateAction31) AddPeriodDetails() *CorporateActionPeriod10 {
	c.PeriodDetails = new(CorporateActionPeriod10)
	return c.PeriodDetails
}

func (c *CorporateAction31) AddRateAndAmountDetails() *CorporateActionRate66 {
	c.RateAndAmountDetails = new(CorporateActionRate66)
	return c.RateAndAmountDetails
}

func (c *CorporateAction31) AddPriceDetails() *CorporateActionPrice57 {
	c.PriceDetails = new(CorporateActionPrice57)
	return c.PriceDetails
}

func (c *CorporateAction31) AddSecuritiesQuantity() *CorporateActionQuantity7 {
	c.SecuritiesQuantity = new(CorporateActionQuantity7)
	return c.SecuritiesQuantity
}

func (c *CorporateAction31) SetInterestAccruedNumberOfDays(value string) {
	c.InterestAccruedNumberOfDays = (*Max3Number)(&value)
}

func (c *CorporateAction31) AddCouponNumber() *IdentificationFormat3Choice {
	newValue := new(IdentificationFormat3Choice)
	c.CouponNumber = append(c.CouponNumber, newValue)
	return newValue
}

func (c *CorporateAction31) SetCertificationBreakdownIndicator(value string) {
	c.CertificationBreakdownIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction31) SetChargesAppliedIndicator(value string) {
	c.ChargesAppliedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction31) SetRestrictionIndicator(value string) {
	c.RestrictionIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction31) SetAccruedInterestIndicator(value string) {
	c.AccruedInterestIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction31) SetLetterOfGuaranteedDeliveryIndicator(value string) {
	c.LetterOfGuaranteedDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction31) AddDividendType() *DividendTypeFormat9Choice {
	c.DividendType = new(DividendTypeFormat9Choice)
	return c.DividendType
}

func (c *CorporateAction31) AddConversionType() *ConversionTypeFormat3Choice {
	c.ConversionType = new(ConversionTypeFormat3Choice)
	return c.ConversionType
}

func (c *CorporateAction31) AddOccurrenceType() *DistributionTypeFormat7Choice {
	c.OccurrenceType = new(DistributionTypeFormat7Choice)
	return c.OccurrenceType
}

func (c *CorporateAction31) AddOfferType() *OfferTypeFormat10Choice {
	newValue := new(OfferTypeFormat10Choice)
	c.OfferType = append(c.OfferType, newValue)
	return newValue
}

func (c *CorporateAction31) AddRenounceableEntitlementStatusType() *RenounceableEntitlementStatusTypeFormat3Choice {
	c.RenounceableEntitlementStatusType = new(RenounceableEntitlementStatusTypeFormat3Choice)
	return c.RenounceableEntitlementStatusType
}

func (c *CorporateAction31) AddEventStage() *CorporateActionEventStageFormat13Choice {
	newValue := new(CorporateActionEventStageFormat13Choice)
	c.EventStage = append(c.EventStage, newValue)
	return newValue
}

func (c *CorporateAction31) AddAdditionalBusinessProcessIndicator() *AdditionalBusinessProcessFormat9Choice {
	newValue := new(AdditionalBusinessProcessFormat9Choice)
	c.AdditionalBusinessProcessIndicator = append(c.AdditionalBusinessProcessIndicator, newValue)
	return newValue
}

func (c *CorporateAction31) AddChangeType() *CorporateActionChangeTypeFormat5Choice {
	newValue := new(CorporateActionChangeTypeFormat5Choice)
	c.ChangeType = append(c.ChangeType, newValue)
	return newValue
}

func (c *CorporateAction31) AddIntermediateSecuritiesDistributionType() *IntermediateSecuritiesDistributionTypeFormat15Choice {
	c.IntermediateSecuritiesDistributionType = new(IntermediateSecuritiesDistributionTypeFormat15Choice)
	return c.IntermediateSecuritiesDistributionType
}

func (c *CorporateAction31) AddCapitalGainInOutIndicator() *CapitalGainFormat3Choice {
	c.CapitalGainInOutIndicator = new(CapitalGainFormat3Choice)
	return c.CapitalGainInOutIndicator
}

func (c *CorporateAction31) AddTaxableIncomePerShareCalculated() *TaxableIncomePerShareCalculatedFormat3Choice {
	c.TaxableIncomePerShareCalculated = new(TaxableIncomePerShareCalculatedFormat3Choice)
	return c.TaxableIncomePerShareCalculated
}

func (c *CorporateAction31) AddElectionType() *ElectionTypeFormat3Choice {
	c.ElectionType = new(ElectionTypeFormat3Choice)
	return c.ElectionType
}

func (c *CorporateAction31) AddLotteryType() *LotteryTypeFormat4Choice {
	c.LotteryType = new(LotteryTypeFormat4Choice)
	return c.LotteryType
}

func (c *CorporateAction31) AddCertificationType() *CertificationTypeFormat3Choice {
	c.CertificationType = new(CertificationTypeFormat3Choice)
	return c.CertificationType
}

func (c *CorporateAction31) AddConsentType() *ConsentTypeFormat4Choice {
	c.ConsentType = new(ConsentTypeFormat4Choice)
	return c.ConsentType
}

func (c *CorporateAction31) AddInformationType() *InformationTypeFormat4Choice {
	c.InformationType = new(InformationTypeFormat4Choice)
	return c.InformationType
}

func (c *CorporateAction31) SetNewPlaceOfIncorporation(value string) {
	c.NewPlaceOfIncorporation = (*Max350Text)(&value)
}

func (c *CorporateAction31) AddAdditionalInformation() *CorporateActionNarrative26 {
	c.AdditionalInformation = new(CorporateActionNarrative26)
	return c.AdditionalInformation
}
