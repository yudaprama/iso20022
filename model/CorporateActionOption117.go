package model

// Provides information about the corporate action option.
type CorporateActionOption117 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption19Choice `xml:"OptnTp"`

	// Specifies the features that may apply to a corporate action option.
	OptionFeatures []*OptionFeaturesFormat18Choice `xml:"OptnFeatrs,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType27Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Provides information about the dates related to a corporate action option.
	DateDetails *CorporateActionDate46 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action option.
	PeriodDetails *CorporateActionPeriod11 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate70 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice61 `xml:"PricDtls,omitempty"`

	// Place where the trade was executed.
	PlaceOfTrade *MarketIdentification84 `xml:"PlcOfTrad,omitempty"`

	// Provides information about the securities movement linked to the corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption50 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement related to a corporate action option.
	CashMovementDetails []*CashOption44 `xml:"CshMvmntDtls,omitempty"`
}

func (c *CorporateActionOption117) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption117) AddOptionType() *CorporateActionOption19Choice {
	c.OptionType = new(CorporateActionOption19Choice)
	return c.OptionType
}

func (c *CorporateActionOption117) AddOptionFeatures() *OptionFeaturesFormat18Choice {
	newValue := new(OptionFeaturesFormat18Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption117) AddFractionDisposition() *FractionDispositionType27Choice {
	c.FractionDisposition = new(FractionDispositionType27Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption117) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption117) AddDateDetails() *CorporateActionDate46 {
	c.DateDetails = new(CorporateActionDate46)
	return c.DateDetails
}

func (c *CorporateActionOption117) AddPeriodDetails() *CorporateActionPeriod11 {
	c.PeriodDetails = new(CorporateActionPeriod11)
	return c.PeriodDetails
}

func (c *CorporateActionOption117) AddRateAndAmountDetails() *CorporateActionRate70 {
	c.RateAndAmountDetails = new(CorporateActionRate70)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption117) AddPriceDetails() *CorporateActionPrice61 {
	c.PriceDetails = new(CorporateActionPrice61)
	return c.PriceDetails
}

func (c *CorporateActionOption117) AddPlaceOfTrade() *MarketIdentification84 {
	c.PlaceOfTrade = new(MarketIdentification84)
	return c.PlaceOfTrade
}

func (c *CorporateActionOption117) AddSecuritiesMovementDetails() *SecuritiesOption50 {
	newValue := new(SecuritiesOption50)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption117) AddCashMovementDetails() *CashOption44 {
	newValue := new(CashOption44)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}
