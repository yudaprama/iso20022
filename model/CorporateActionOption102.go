package model

// Provides information about the corporate action option.
type CorporateActionOption102 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption11Choice `xml:"OptnTp"`

	// Specifies the features that may apply to a corporate action option.
	OptionFeatures []*OptionFeaturesFormat1Choice `xml:"OptnFeatrs,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType23Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Provides information about the dates related to a corporate action option.
	DateDetails *CorporateActionDate18 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action option.
	PeriodDetails *CorporateActionPeriod9 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate46 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice30 `xml:"PricDtls,omitempty"`

	// Place where the trade was executed.
	PlaceOfTrade *MarketIdentification78 `xml:"PlcOfTrad,omitempty"`

	// Provides information about the securities movement linked to the corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption42 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement related to a corporate action option.
	CashMovementDetails []*CashOption30 `xml:"CshMvmntDtls,omitempty"`
}

func (c *CorporateActionOption102) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption102) AddOptionType() *CorporateActionOption11Choice {
	c.OptionType = new(CorporateActionOption11Choice)
	return c.OptionType
}

func (c *CorporateActionOption102) AddOptionFeatures() *OptionFeaturesFormat1Choice {
	newValue := new(OptionFeaturesFormat1Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption102) AddFractionDisposition() *FractionDispositionType23Choice {
	c.FractionDisposition = new(FractionDispositionType23Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption102) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption102) AddDateDetails() *CorporateActionDate18 {
	c.DateDetails = new(CorporateActionDate18)
	return c.DateDetails
}

func (c *CorporateActionOption102) AddPeriodDetails() *CorporateActionPeriod9 {
	c.PeriodDetails = new(CorporateActionPeriod9)
	return c.PeriodDetails
}

func (c *CorporateActionOption102) AddRateAndAmountDetails() *CorporateActionRate46 {
	c.RateAndAmountDetails = new(CorporateActionRate46)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption102) AddPriceDetails() *CorporateActionPrice30 {
	c.PriceDetails = new(CorporateActionPrice30)
	return c.PriceDetails
}

func (c *CorporateActionOption102) AddPlaceOfTrade() *MarketIdentification78 {
	c.PlaceOfTrade = new(MarketIdentification78)
	return c.PlaceOfTrade
}

func (c *CorporateActionOption102) AddSecuritiesMovementDetails() *SecuritiesOption42 {
	newValue := new(SecuritiesOption42)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption102) AddCashMovementDetails() *CashOption30 {
	newValue := new(CashOption30)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}
