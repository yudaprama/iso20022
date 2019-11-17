package model

// Provides information about the corporate action option.
type CorporateActionOption124 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption26Choice `xml:"OptnTp"`

	// Specifies the features that may apply to a corporate action option.
	OptionFeatures []*OptionFeaturesFormat19Choice `xml:"OptnFeatrs,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType30Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Provides information about the dates related to a corporate action option.
	DateDetails *CorporateActionDate52 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action option.
	PeriodDetails *CorporateActionPeriod11 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate74 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice63 `xml:"PricDtls,omitempty"`

	// Place where the trade was executed.
	PlaceOfTrade *MarketIdentification90 `xml:"PlcOfTrad,omitempty"`

	// Provides information about the securities movement linked to the corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption55 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement related to a corporate action option.
	CashMovementDetails []*CashOption46 `xml:"CshMvmntDtls,omitempty"`
}

func (c *CorporateActionOption124) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption124) AddOptionType() *CorporateActionOption26Choice {
	c.OptionType = new(CorporateActionOption26Choice)
	return c.OptionType
}

func (c *CorporateActionOption124) AddOptionFeatures() *OptionFeaturesFormat19Choice {
	newValue := new(OptionFeaturesFormat19Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption124) AddFractionDisposition() *FractionDispositionType30Choice {
	c.FractionDisposition = new(FractionDispositionType30Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption124) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption124) AddDateDetails() *CorporateActionDate52 {
	c.DateDetails = new(CorporateActionDate52)
	return c.DateDetails
}

func (c *CorporateActionOption124) AddPeriodDetails() *CorporateActionPeriod11 {
	c.PeriodDetails = new(CorporateActionPeriod11)
	return c.PeriodDetails
}

func (c *CorporateActionOption124) AddRateAndAmountDetails() *CorporateActionRate74 {
	c.RateAndAmountDetails = new(CorporateActionRate74)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption124) AddPriceDetails() *CorporateActionPrice63 {
	c.PriceDetails = new(CorporateActionPrice63)
	return c.PriceDetails
}

func (c *CorporateActionOption124) AddPlaceOfTrade() *MarketIdentification90 {
	c.PlaceOfTrade = new(MarketIdentification90)
	return c.PlaceOfTrade
}

func (c *CorporateActionOption124) AddSecuritiesMovementDetails() *SecuritiesOption55 {
	newValue := new(SecuritiesOption55)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption124) AddCashMovementDetails() *CashOption46 {
	newValue := new(CashOption46)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}
