package model

// Provides information about the corporate action option.
type CorporateActionOption4 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption3Choice `xml:"OptnTp"`

	// Specifies the features that may apply to a corporate action option.
	OptionFeatures []*OptionFeaturesFormat1Choice `xml:"OptnFeatrs,omitempty"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType1Choice `xml:"FrctnDspstn,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Provides information about the dates related to a corporate action option.
	DateDetails *CorporateActionDate6 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action option.
	PeriodDetails *CorporateActionPeriod4 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate4 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice5 `xml:"PricDtls,omitempty"`

	// Place where the trade was executed.
	PlaceOfTrade *MarketIdentification4 `xml:"PlcOfTrad,omitempty"`

	// Provides information about the securities movement linked to the corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption3 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement related to a corporate action option.
	CashMovementDetails []*CashOption2 `xml:"CshMvmntDtls,omitempty"`
}

func (c *CorporateActionOption4) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption4) AddOptionType() *CorporateActionOption3Choice {
	c.OptionType = new(CorporateActionOption3Choice)
	return c.OptionType
}

func (c *CorporateActionOption4) AddOptionFeatures() *OptionFeaturesFormat1Choice {
	newValue := new(OptionFeaturesFormat1Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption4) AddFractionDisposition() *FractionDispositionType1Choice {
	c.FractionDisposition = new(FractionDispositionType1Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption4) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption4) AddDateDetails() *CorporateActionDate6 {
	c.DateDetails = new(CorporateActionDate6)
	return c.DateDetails
}

func (c *CorporateActionOption4) AddPeriodDetails() *CorporateActionPeriod4 {
	c.PeriodDetails = new(CorporateActionPeriod4)
	return c.PeriodDetails
}

func (c *CorporateActionOption4) AddRateAndAmountDetails() *CorporateActionRate4 {
	c.RateAndAmountDetails = new(CorporateActionRate4)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption4) AddPriceDetails() *CorporateActionPrice5 {
	c.PriceDetails = new(CorporateActionPrice5)
	return c.PriceDetails
}

func (c *CorporateActionOption4) AddPlaceOfTrade() *MarketIdentification4 {
	c.PlaceOfTrade = new(MarketIdentification4)
	return c.PlaceOfTrade
}

func (c *CorporateActionOption4) AddSecuritiesMovementDetails() *SecuritiesOption3 {
	newValue := new(SecuritiesOption3)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption4) AddCashMovementDetails() *CashOption2 {
	newValue := new(CashOption2)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}
