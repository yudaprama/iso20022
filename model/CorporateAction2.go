package model

// Provides information about the CA event.
type CorporateAction2 struct {

	// Stage in the corporate action event life cycle.
	EventStage []*CorporateActionEventStage1FormatChoice `xml:"EvtStag,omitempty"`

	// Identifies the option that will be selected by default if no instruction is provided by account owner.
	DefaultOptionType *CorporateActionOption1FormatChoice `xml:"DfltOptnTp,omitempty"`

	// Identifies the option number that will be selected by default if no instruction is provided by account owner.
	DefaultOptionNumber *Exact3NumericText `xml:"DfltOptnNb,omitempty"`

	// The method of calculation of drawings and partial redemptions.
	CalculationMethod *CorporateActionCalculationMethod1FormatChoice `xml:"ClctnMtd,omitempty"`

	// Represents the presence of a back end odd lot provision and the quantity of equity required after proration to be eligible for this privilege.
	BackEndOddLotSecuritiesQuantity *UnitOrFaceAmountOrCode1Choice `xml:"BckEndOddLotSctiesQty,omitempty"`

	// Specifies that if an order is prorated holders of odd lots who tender their full position will not have tendered position prorated but rather accepted in full.
	FrontEndOddLotSecuritiesQuantity *UnitOrFaceAmountOrCode1Choice `xml:"FrntEndOddLotSctiesQty,omitempty"`

	// Minimum quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableSecuritiesQuantity *UnitOrFaceAmount1Choice `xml:"MinExrcblSctiesQty,omitempty"`

	// Minimum multiple quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableMultipleSecuritiesQuantity *UnitOrFaceAmount1Choice `xml:"MinExrcblMltplSctiesQty,omitempty"`

	// Amount used when the called amount is not met by running the lottery with the base denomination.
	IncrementalDenomination *UnitOrFaceAmount1Choice `xml:"IncrmtlDnmtn,omitempty"`

	// New Denomination of the equity following, eg, an increase or decrease in nominal value.
	NewDenominationSecuritiesQuantity *UnitOrFaceAmount1Choice `xml:"NewDnmtnSctiesQty,omitempty"`

	// Quantity of equity that makes up the new board lot.
	NewBoardLotSecuritiesQuantity *UnitOrFaceAmount1Choice `xml:"NewBrdLotSctiesQty,omitempty"`

	// Quantity of securities the offeror/issuer will purchase or redeem under the terms of the event. This can be a number or the term "any and all".
	SecuritiesQuantitySought *UnitOrFaceAmountOrCode1Choice `xml:"SctiesQtySght,omitempty"`

	// The minimum integral amount of securities that each account owner must have remaining after the called amounts are applied.
	BaseDenomination *UnitOrFaceAmount1Choice `xml:"BaseDnmtn,omitempty"`

	// Specifies the type of change announced.
	ChangeType []*CorporateActionChangeType1FormatChoice `xml:"ChngTp,omitempty"`

	// Specifies the conditions that apply to the offer.
	OfferType []*OfferType1FormatChoice `xml:"OfferTp,omitempty"`

	// Indicates whether there is restrictions apply to the CA event or not.
	//
	// Yes = There is restrictions.
	// No = There is no restrictions.
	RestrictionIndicator *YesNoIndicator `xml:"RstrctnInd,omitempty"`

	// Specifies if the issuer will allow the agent to accept partial elections. It is to allow split voting over options. It allows the client to elect more than one option to be selected per designated holding.
	PartialElectionIndicator *YesNoIndicator `xml:"PrtlElctnInd,omitempty"`

	// Specifies the effect on the holdings of electing a Corporate Action option.
	ElectionType *ElectionMovementType1FormatChoice `xml:"ElctnTp,omitempty"`

	// Specifies the type of lottery announced.
	LotteryType *LotteryType1FormatChoice `xml:"LtryTp,omitempty"`

	// Specifies the type of income.
	// The lists of income type codes to be used, are available on the SMPG website at www.smpg.info.
	IncomeType *GenericIdentification13 `xml:"IncmTp,omitempty"`

	// Specifies the conditions in which a dividend is paid.
	DividendType *CorporateActionFrequencyType1FormatChoice `xml:"DvddTp,omitempty"`

	// Type of intermediates securities distribution, eg, stock dividend, reverse right.
	IntermediateSecuritiesDistributionType *IntermediateSecurityDistributionType1FormatChoice `xml:"IntrmdtSctiesDstrbtnTp,omitempty"`

	// Number of the coupon attached/associated with a security.
	CouponNumber []*Max3NumericText `xml:"CpnNb,omitempty"`

	// Number of days used for calculating the accrued interest amount.
	InterestAccruedNumberOfDays *Number `xml:"IntrstAcrdNbOfDays,omitempty"`

	// New denomination currency of the inancial instrument.
	NewDenominationCurrency *ActiveCurrencyCode `xml:"NewDnmtnCcy,omitempty"`

	// Provides information about the dates related to a CA event.
	DateDetails *CorporateActionDate2 `xml:"DtDtls,omitempty"`

	// Provides information about the prices related to a CA event.
	PriceDetails []*CorporateActionPrice2 `xml:"PricDtls,omitempty"`

	// Provides information about the periods related to a CA event.
	PeriodDetails *CorporateActionPeriod1 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a CA event.
	RateAndAmountDetails *CorporateActionRate1 `xml:"RateAndAmtDtls,omitempty"`

	// Provides additional information.
	CorporateActionAdditionalInformation *CorporateActionNarrative1 `xml:"CorpActnAddtlInf,omitempty"`

	// Indicates whether certification is required from the account owner.
	CertificationRequiredIndicator *YesNoIndicator `xml:"CertfctnReqrdInd,omitempty"`

	// Type of certification which is required.
	CertificationType *BeneficiaryCertificationType1FormatChoice `xml:"CertfctnTp,omitempty"`

	// Specifies whether the capital gain is in the scope of the EU Savings directive for the income realised upon the sale, refund or redemption of shares and units (...) (Article 6(1d)).
	CapitalGain *EUCapitalGain2Code `xml:"CptlGn,omitempty"`

	// Specifies whether the financial instrument calculates the taxable income per dividend/taxable income per share.
	TaxableIncomePerShareCalculated *TaxableIncomePerShareCalculated2Code `xml:"TaxblIncmPerShrClctd,omitempty"`

	// New companys place of incorporation.
	NewPlaceOfIncorporation *Max70Text `xml:"NewPlcOfIncorprtn,omitempty"`

	// Specifies whether terms of the event allow resale of the rights.
	RenounceableEntitlementStatusType *RenounceableStatus1FormatChoice `xml:"RnncblEntitlmntStsTp,omitempty"`

	// Specifies the conversion type of an instrument.
	ConversionType *ConversionType1FormatChoice `xml:"ConvsTp,omitempty"`

	// Indicates whether redemption charges apply.
	RedemptionChargesAppliedIndicator *YesNoIndicator `xml:"RedChrgsApldInd,omitempty"`

	// Specifies whether the proceeds of the event will be distributed on a rolling basis rather than on a specific date.
	DistributionType *DistributionType1FormatChoice `xml:"DstrbtnTp,omitempty"`
}

func (c *CorporateAction2) AddEventStage() *CorporateActionEventStage1FormatChoice {
	newValue := new(CorporateActionEventStage1FormatChoice)
	c.EventStage = append(c.EventStage, newValue)
	return newValue
}

func (c *CorporateAction2) AddDefaultOptionType() *CorporateActionOption1FormatChoice {
	c.DefaultOptionType = new(CorporateActionOption1FormatChoice)
	return c.DefaultOptionType
}

func (c *CorporateAction2) SetDefaultOptionNumber(value string) {
	c.DefaultOptionNumber = (*Exact3NumericText)(&value)
}

func (c *CorporateAction2) AddCalculationMethod() *CorporateActionCalculationMethod1FormatChoice {
	c.CalculationMethod = new(CorporateActionCalculationMethod1FormatChoice)
	return c.CalculationMethod
}

func (c *CorporateAction2) AddBackEndOddLotSecuritiesQuantity() *UnitOrFaceAmountOrCode1Choice {
	c.BackEndOddLotSecuritiesQuantity = new(UnitOrFaceAmountOrCode1Choice)
	return c.BackEndOddLotSecuritiesQuantity
}

func (c *CorporateAction2) AddFrontEndOddLotSecuritiesQuantity() *UnitOrFaceAmountOrCode1Choice {
	c.FrontEndOddLotSecuritiesQuantity = new(UnitOrFaceAmountOrCode1Choice)
	return c.FrontEndOddLotSecuritiesQuantity
}

func (c *CorporateAction2) AddMinimumExercisableSecuritiesQuantity() *UnitOrFaceAmount1Choice {
	c.MinimumExercisableSecuritiesQuantity = new(UnitOrFaceAmount1Choice)
	return c.MinimumExercisableSecuritiesQuantity
}

func (c *CorporateAction2) AddMinimumExercisableMultipleSecuritiesQuantity() *UnitOrFaceAmount1Choice {
	c.MinimumExercisableMultipleSecuritiesQuantity = new(UnitOrFaceAmount1Choice)
	return c.MinimumExercisableMultipleSecuritiesQuantity
}

func (c *CorporateAction2) AddIncrementalDenomination() *UnitOrFaceAmount1Choice {
	c.IncrementalDenomination = new(UnitOrFaceAmount1Choice)
	return c.IncrementalDenomination
}

func (c *CorporateAction2) AddNewDenominationSecuritiesQuantity() *UnitOrFaceAmount1Choice {
	c.NewDenominationSecuritiesQuantity = new(UnitOrFaceAmount1Choice)
	return c.NewDenominationSecuritiesQuantity
}

func (c *CorporateAction2) AddNewBoardLotSecuritiesQuantity() *UnitOrFaceAmount1Choice {
	c.NewBoardLotSecuritiesQuantity = new(UnitOrFaceAmount1Choice)
	return c.NewBoardLotSecuritiesQuantity
}

func (c *CorporateAction2) AddSecuritiesQuantitySought() *UnitOrFaceAmountOrCode1Choice {
	c.SecuritiesQuantitySought = new(UnitOrFaceAmountOrCode1Choice)
	return c.SecuritiesQuantitySought
}

func (c *CorporateAction2) AddBaseDenomination() *UnitOrFaceAmount1Choice {
	c.BaseDenomination = new(UnitOrFaceAmount1Choice)
	return c.BaseDenomination
}

func (c *CorporateAction2) AddChangeType() *CorporateActionChangeType1FormatChoice {
	newValue := new(CorporateActionChangeType1FormatChoice)
	c.ChangeType = append(c.ChangeType, newValue)
	return newValue
}

func (c *CorporateAction2) AddOfferType() *OfferType1FormatChoice {
	newValue := new(OfferType1FormatChoice)
	c.OfferType = append(c.OfferType, newValue)
	return newValue
}

func (c *CorporateAction2) SetRestrictionIndicator(value string) {
	c.RestrictionIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction2) SetPartialElectionIndicator(value string) {
	c.PartialElectionIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction2) AddElectionType() *ElectionMovementType1FormatChoice {
	c.ElectionType = new(ElectionMovementType1FormatChoice)
	return c.ElectionType
}

func (c *CorporateAction2) AddLotteryType() *LotteryType1FormatChoice {
	c.LotteryType = new(LotteryType1FormatChoice)
	return c.LotteryType
}

func (c *CorporateAction2) AddIncomeType() *GenericIdentification13 {
	c.IncomeType = new(GenericIdentification13)
	return c.IncomeType
}

func (c *CorporateAction2) AddDividendType() *CorporateActionFrequencyType1FormatChoice {
	c.DividendType = new(CorporateActionFrequencyType1FormatChoice)
	return c.DividendType
}

func (c *CorporateAction2) AddIntermediateSecuritiesDistributionType() *IntermediateSecurityDistributionType1FormatChoice {
	c.IntermediateSecuritiesDistributionType = new(IntermediateSecurityDistributionType1FormatChoice)
	return c.IntermediateSecuritiesDistributionType
}

func (c *CorporateAction2) AddCouponNumber(value string) {
	c.CouponNumber = append(c.CouponNumber, (*Max3NumericText)(&value))
}

func (c *CorporateAction2) SetInterestAccruedNumberOfDays(value string) {
	c.InterestAccruedNumberOfDays = (*Number)(&value)
}

func (c *CorporateAction2) SetNewDenominationCurrency(value string) {
	c.NewDenominationCurrency = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateAction2) AddDateDetails() *CorporateActionDate2 {
	c.DateDetails = new(CorporateActionDate2)
	return c.DateDetails
}

func (c *CorporateAction2) AddPriceDetails() *CorporateActionPrice2 {
	newValue := new(CorporateActionPrice2)
	c.PriceDetails = append(c.PriceDetails, newValue)
	return newValue
}

func (c *CorporateAction2) AddPeriodDetails() *CorporateActionPeriod1 {
	c.PeriodDetails = new(CorporateActionPeriod1)
	return c.PeriodDetails
}

func (c *CorporateAction2) AddRateAndAmountDetails() *CorporateActionRate1 {
	c.RateAndAmountDetails = new(CorporateActionRate1)
	return c.RateAndAmountDetails
}

func (c *CorporateAction2) AddCorporateActionAdditionalInformation() *CorporateActionNarrative1 {
	c.CorporateActionAdditionalInformation = new(CorporateActionNarrative1)
	return c.CorporateActionAdditionalInformation
}

func (c *CorporateAction2) SetCertificationRequiredIndicator(value string) {
	c.CertificationRequiredIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction2) AddCertificationType() *BeneficiaryCertificationType1FormatChoice {
	c.CertificationType = new(BeneficiaryCertificationType1FormatChoice)
	return c.CertificationType
}

func (c *CorporateAction2) SetCapitalGain(value string) {
	c.CapitalGain = (*EUCapitalGain2Code)(&value)
}

func (c *CorporateAction2) SetTaxableIncomePerShareCalculated(value string) {
	c.TaxableIncomePerShareCalculated = (*TaxableIncomePerShareCalculated2Code)(&value)
}

func (c *CorporateAction2) SetNewPlaceOfIncorporation(value string) {
	c.NewPlaceOfIncorporation = (*Max70Text)(&value)
}

func (c *CorporateAction2) AddRenounceableEntitlementStatusType() *RenounceableStatus1FormatChoice {
	c.RenounceableEntitlementStatusType = new(RenounceableStatus1FormatChoice)
	return c.RenounceableEntitlementStatusType
}

func (c *CorporateAction2) AddConversionType() *ConversionType1FormatChoice {
	c.ConversionType = new(ConversionType1FormatChoice)
	return c.ConversionType
}

func (c *CorporateAction2) SetRedemptionChargesAppliedIndicator(value string) {
	c.RedemptionChargesAppliedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateAction2) AddDistributionType() *DistributionType1FormatChoice {
	c.DistributionType = new(DistributionType1FormatChoice)
	return c.DistributionType
}
