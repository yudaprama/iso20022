package model

// Provides information about the corporate action option.
type CorporateActionOption132 struct {

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
	RateAndAmountDetails *CorporateActionRate82 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice61 `xml:"PricDtls,omitempty"`

	// Place where the trade was executed.
	PlaceOfTrade *MarketIdentification84 `xml:"PlcOfTrad,omitempty"`

	// Provides information about the securities movement linked to the corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption60 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement related to a corporate action option.
	CashMovementDetails []*CashOption52 `xml:"CshMvmntDtls,omitempty"`
}

func (c *CorporateActionOption132) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption132) AddOptionType() *CorporateActionOption19Choice {
	c.OptionType = new(CorporateActionOption19Choice)
	return c.OptionType
}

func (c *CorporateActionOption132) AddOptionFeatures() *OptionFeaturesFormat18Choice {
	newValue := new(OptionFeaturesFormat18Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption132) AddFractionDisposition() *FractionDispositionType27Choice {
	c.FractionDisposition = new(FractionDispositionType27Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption132) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption132) AddDateDetails() *CorporateActionDate46 {
	c.DateDetails = new(CorporateActionDate46)
	return c.DateDetails
}

func (c *CorporateActionOption132) AddPeriodDetails() *CorporateActionPeriod11 {
	c.PeriodDetails = new(CorporateActionPeriod11)
	return c.PeriodDetails
}

func (c *CorporateActionOption132) AddRateAndAmountDetails() *CorporateActionRate82 {
	c.RateAndAmountDetails = new(CorporateActionRate82)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption132) AddPriceDetails() *CorporateActionPrice61 {
	c.PriceDetails = new(CorporateActionPrice61)
	return c.PriceDetails
}

func (c *CorporateActionOption132) AddPlaceOfTrade() *MarketIdentification84 {
	c.PlaceOfTrade = new(MarketIdentification84)
	return c.PlaceOfTrade
}

func (c *CorporateActionOption132) AddSecuritiesMovementDetails() *SecuritiesOption60 {
	newValue := new(SecuritiesOption60)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption132) AddCashMovementDetails() *CashOption52 {
	newValue := new(CashOption52)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}
