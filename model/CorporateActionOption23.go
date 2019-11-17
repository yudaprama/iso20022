package model

// Provides information about the corporate action option.
type CorporateActionOption23 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption3Choice `xml:"OptnTp"`

	// Specifies the features that may apply to a corporate action option.
	OptionFeatures []*OptionFeaturesFormat1Choice `xml:"OptnFeatrs,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType12Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Provides information about the dates related to a corporate action option.
	DateDetails *CorporateActionDate18 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action option.
	PeriodDetails *CorporateActionPeriod9 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate20 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice21 `xml:"PricDtls,omitempty"`

	// Place where the trade was executed.
	PlaceOfTrade *MarketIdentification4 `xml:"PlcOfTrad,omitempty"`

	// Provides information about the securities movement linked to the corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption18 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement related to a corporate action option.
	CashMovementDetails []*CashOption12 `xml:"CshMvmntDtls,omitempty"`
}

func (c *CorporateActionOption23) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption23) AddOptionType() *CorporateActionOption3Choice {
	c.OptionType = new(CorporateActionOption3Choice)
	return c.OptionType
}

func (c *CorporateActionOption23) AddOptionFeatures() *OptionFeaturesFormat1Choice {
	newValue := new(OptionFeaturesFormat1Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption23) AddFractionDisposition() *FractionDispositionType12Choice {
	c.FractionDisposition = new(FractionDispositionType12Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption23) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption23) AddDateDetails() *CorporateActionDate18 {
	c.DateDetails = new(CorporateActionDate18)
	return c.DateDetails
}

func (c *CorporateActionOption23) AddPeriodDetails() *CorporateActionPeriod9 {
	c.PeriodDetails = new(CorporateActionPeriod9)
	return c.PeriodDetails
}

func (c *CorporateActionOption23) AddRateAndAmountDetails() *CorporateActionRate20 {
	c.RateAndAmountDetails = new(CorporateActionRate20)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption23) AddPriceDetails() *CorporateActionPrice21 {
	c.PriceDetails = new(CorporateActionPrice21)
	return c.PriceDetails
}

func (c *CorporateActionOption23) AddPlaceOfTrade() *MarketIdentification4 {
	c.PlaceOfTrade = new(MarketIdentification4)
	return c.PlaceOfTrade
}

func (c *CorporateActionOption23) AddSecuritiesMovementDetails() *SecuritiesOption18 {
	newValue := new(SecuritiesOption18)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption23) AddCashMovementDetails() *CashOption12 {
	newValue := new(CashOption12)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}
