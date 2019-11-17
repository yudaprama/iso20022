package model

// Provides information about the corporate action option.
type CorporateActionOption136 struct {

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
	RateAndAmountDetails *CorporateActionRate85 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice63 `xml:"PricDtls,omitempty"`

	// Place where the trade was executed.
	PlaceOfTrade *MarketIdentification90 `xml:"PlcOfTrad,omitempty"`

	// Provides information about the securities movement linked to the corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption63 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement related to a corporate action option.
	CashMovementDetails []*CashOption55 `xml:"CshMvmntDtls,omitempty"`
}

func (c *CorporateActionOption136) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption136) AddOptionType() *CorporateActionOption26Choice {
	c.OptionType = new(CorporateActionOption26Choice)
	return c.OptionType
}

func (c *CorporateActionOption136) AddOptionFeatures() *OptionFeaturesFormat19Choice {
	newValue := new(OptionFeaturesFormat19Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption136) AddFractionDisposition() *FractionDispositionType30Choice {
	c.FractionDisposition = new(FractionDispositionType30Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption136) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption136) AddDateDetails() *CorporateActionDate52 {
	c.DateDetails = new(CorporateActionDate52)
	return c.DateDetails
}

func (c *CorporateActionOption136) AddPeriodDetails() *CorporateActionPeriod11 {
	c.PeriodDetails = new(CorporateActionPeriod11)
	return c.PeriodDetails
}

func (c *CorporateActionOption136) AddRateAndAmountDetails() *CorporateActionRate85 {
	c.RateAndAmountDetails = new(CorporateActionRate85)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption136) AddPriceDetails() *CorporateActionPrice63 {
	c.PriceDetails = new(CorporateActionPrice63)
	return c.PriceDetails
}

func (c *CorporateActionOption136) AddPlaceOfTrade() *MarketIdentification90 {
	c.PlaceOfTrade = new(MarketIdentification90)
	return c.PlaceOfTrade
}

func (c *CorporateActionOption136) AddSecuritiesMovementDetails() *SecuritiesOption63 {
	newValue := new(SecuritiesOption63)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption136) AddCashMovementDetails() *CashOption55 {
	newValue := new(CashOption55)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}
