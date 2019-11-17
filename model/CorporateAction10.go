package model

// Provides information about the corporate action event.
type CorporateAction10 struct {

	// Provides information about the dates related to a corporate action event.
	DateDetails *CorporateActionDate27 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action event.
	PeriodDetails *CorporateActionPeriod8 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action event.
	RateAndAmountDetails *CorporateActionRate35 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action event.
	PriceDetails *CorporateActionPrice17 `xml:"PricDtls,omitempty"`

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
	DividendType *DividendTypeFormat3Choice `xml:"DvddTp,omitempty"`

	// Specifies the conversion type of an instrument.
	ConversionType *ConversionTypeFormat1Choice `xml:"ConvsTp,omitempty"`

	// Specifies the conditions in which the payment of the proceeds occurs.
	PaymentOccurrenceType *DistributionTypeFormat3Choice `xml:"PmtOcrncTp,omitempty"`

	// Specifies the conditions that apply to the offer.
	OfferType []*OfferTypeFormat3Choice `xml:"OfferTp,omitempty"`

	// Specifies whether terms of the event allow resale of the rights.
	RenounceableEntitlementStatusType *RenounceableEntitlementStatusTypeFormat1Choice `xml:"RnncblEntitlmntStsTp,omitempty"`

	// Stage in the corporate action event life cycle.
	EventStage []*CorporateActionEventStageFormat3Choice `xml:"EvtStag,omitempty"`

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

	// New company's place of incorporation.
	NewPlaceOfIncorporation *Max350Text `xml:"NewPlcOfIncorprtn,omitempty"`

	// Provides additional information. This field may only be used when the information to be transmitted, cannot be coded.
	AdditionalInformation *CorporateActionNarrative24 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateAction10) AddDateDetails() *CorporateActionDate27 {
	c.DateDetails = new(CorporateActionDate27)
	return c.DateDetails
}

func (c *CorporateAction10) AddPeriodDetails() *CorporateActionPeriod8 {
	c.PeriodDetails = new(CorporateActionPeriod8)
	return c.PeriodDetails
}

func (c *CorporateAction10) AddRateAndAmountDetails() *CorporateActionRate35 {
	c.RateAndAmountDetails = new(CorporateActionRate35)
	return c.RateAndAmountDetails
}

func (c *CorporateAction10) AddPriceDetails() *CorporateActionPrice17 {
	c.PriceDetails = new(CorporateActionPrice17)
	return c.PriceDetails
}

func (c *CorporateAction10) AddSecuritiesQuantity() *CorporateActionQuantity5 {
	c.SecuritiesQuantity = new(CorporateActionQuantity5)
	return c.SecuritiesQuantity
}

func (c *CorporateAction10) SetInterestAccruedNumberOfDays(value string) {
	c.InterestAccruedNumberOfDays = (*Max3Number)(&value)
}

func (c *CorporateAction10) AddCouponNumber() *IdentificationFormat1Choice {
	newValue := new(IdentificationFormat1Choice)
	c.CouponNumber = append(c.CouponNumber, newValue)
	return newValue
}

func (c *CorporateAction10) SetCertificationBreakdownIndicator(value string) {
	c.CertificationBreakdownIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction10) SetChargesAppliedIndicator(value string) {
	c.ChargesAppliedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction10) SetRestrictionIndicator(value string) {
	c.RestrictionIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction10) SetAccruedInterestIndicator(value string) {
	c.AccruedInterestIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction10) SetLetterOfGuaranteedDeliveryIndicator(value string) {
	c.LetterOfGuaranteedDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction10) AddDividendType() *DividendTypeFormat3Choice {
	c.DividendType = new(DividendTypeFormat3Choice)
	return c.DividendType
}

func (c *CorporateAction10) AddConversionType() *ConversionTypeFormat1Choice {
	c.ConversionType = new(ConversionTypeFormat1Choice)
	return c.ConversionType
}

func (c *CorporateAction10) AddPaymentOccurrenceType() *DistributionTypeFormat3Choice {
	c.PaymentOccurrenceType = new(DistributionTypeFormat3Choice)
	return c.PaymentOccurrenceType
}

func (c *CorporateAction10) AddOfferType() *OfferTypeFormat3Choice {
	newValue := new(OfferTypeFormat3Choice)
	c.OfferType = append(c.OfferType, newValue)
	return newValue
}

func (c *CorporateAction10) AddRenounceableEntitlementStatusType() *RenounceableEntitlementStatusTypeFormat1Choice {
	c.RenounceableEntitlementStatusType = new(RenounceableEntitlementStatusTypeFormat1Choice)
	return c.RenounceableEntitlementStatusType
}

func (c *CorporateAction10) AddEventStage() *CorporateActionEventStageFormat3Choice {
	newValue := new(CorporateActionEventStageFormat3Choice)
	c.EventStage = append(c.EventStage, newValue)
	return newValue
}

func (c *CorporateAction10) AddAdditionalBusinessProcessIndicator() *AdditionalBusinessProcessFormat1Choice {
	newValue := new(AdditionalBusinessProcessFormat1Choice)
	c.AdditionalBusinessProcessIndicator = append(c.AdditionalBusinessProcessIndicator, newValue)
	return newValue
}

func (c *CorporateAction10) AddChangeType() *CorporateActionChangeTypeFormat1Choice {
	newValue := new(CorporateActionChangeTypeFormat1Choice)
	c.ChangeType = append(c.ChangeType, newValue)
	return newValue
}

func (c *CorporateAction10) AddIntermediateSecuritiesDistributionType() *IntermediateSecuritiesDistributionTypeFormat9Choice {
	c.IntermediateSecuritiesDistributionType = new(IntermediateSecuritiesDistributionTypeFormat9Choice)
	return c.IntermediateSecuritiesDistributionType
}

func (c *CorporateAction10) AddCapitalGainInOutIndicator() *CapitalGainFormat1Choice {
	c.CapitalGainInOutIndicator = new(CapitalGainFormat1Choice)
	return c.CapitalGainInOutIndicator
}

func (c *CorporateAction10) AddTaxableIncomePerShareCalculated() *TaxableIncomePerShareCalculatedFormat1Choice {
	c.TaxableIncomePerShareCalculated = new(TaxableIncomePerShareCalculatedFormat1Choice)
	return c.TaxableIncomePerShareCalculated
}

func (c *CorporateAction10) AddElectionType() *ElectionTypeFormat1Choice {
	c.ElectionType = new(ElectionTypeFormat1Choice)
	return c.ElectionType
}

func (c *CorporateAction10) AddLotteryType() *LotteryTypeFormat1Choice {
	c.LotteryType = new(LotteryTypeFormat1Choice)
	return c.LotteryType
}

func (c *CorporateAction10) AddCertificationType() *CertificationTypeFormat1Choice {
	c.CertificationType = new(CertificationTypeFormat1Choice)
	return c.CertificationType
}

func (c *CorporateAction10) SetNewPlaceOfIncorporation(value string) {
	c.NewPlaceOfIncorporation = (*Max350Text)(&value)
}

func (c *CorporateAction10) AddAdditionalInformation() *CorporateActionNarrative24 {
	c.AdditionalInformation = new(CorporateActionNarrative24)
	return c.AdditionalInformation
}
