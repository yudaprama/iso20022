package model

// Provides information about the corporate action option.
type CorporateActionOption53 struct {

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
	RateAndAmountDetails *CorporateActionRate38 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice30 `xml:"PricDtls,omitempty"`

	// Place where the trade was executed.
	PlaceOfTrade *MarketIdentification4 `xml:"PlcOfTrad,omitempty"`

	// Provides information about the securities movement linked to the corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption35 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement related to a corporate action option.
	CashMovementDetails []*CashOption26 `xml:"CshMvmntDtls,omitempty"`
}

func (c *CorporateActionOption53) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption53) AddOptionType() *CorporateActionOption11Choice {
	c.OptionType = new(CorporateActionOption11Choice)
	return c.OptionType
}

func (c *CorporateActionOption53) AddOptionFeatures() *OptionFeaturesFormat1Choice {
	newValue := new(OptionFeaturesFormat1Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption53) AddFractionDisposition() *FractionDispositionType23Choice {
	c.FractionDisposition = new(FractionDispositionType23Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption53) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption53) AddDateDetails() *CorporateActionDate18 {
	c.DateDetails = new(CorporateActionDate18)
	return c.DateDetails
}

func (c *CorporateActionOption53) AddPeriodDetails() *CorporateActionPeriod9 {
	c.PeriodDetails = new(CorporateActionPeriod9)
	return c.PeriodDetails
}

func (c *CorporateActionOption53) AddRateAndAmountDetails() *CorporateActionRate38 {
	c.RateAndAmountDetails = new(CorporateActionRate38)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption53) AddPriceDetails() *CorporateActionPrice30 {
	c.PriceDetails = new(CorporateActionPrice30)
	return c.PriceDetails
}

func (c *CorporateActionOption53) AddPlaceOfTrade() *MarketIdentification4 {
	c.PlaceOfTrade = new(MarketIdentification4)
	return c.PlaceOfTrade
}

func (c *CorporateActionOption53) AddSecuritiesMovementDetails() *SecuritiesOption35 {
	newValue := new(SecuritiesOption35)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption53) AddCashMovementDetails() *CashOption26 {
	newValue := new(CashOption26)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}
