package model

// Provides information about the corporate action option.
type CorporateActionOption131 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption20Choice `xml:"OptnTp"`

	// Specifies how fractional amount/quantities are treated.
	FractionDisposition *FractionDispositionType28Choice `xml:"FrctnDspstn,omitempty"`

	// Type of changes affecting the security form.
	ChangeType []*CorporateActionChangeTypeFormat6Choice `xml:"ChngTp,omitempty"`

	// Specifies that the corporate action instruction is to be processed using the Available-for-Collateral pool.
	EligibleForCollateralIndicator *YesNoIndicator `xml:"ElgblForCollInd,omitempty"`

	// Account servicer is instructed to buy the indicated currency after the receipt of cash proceeds.
	CurrencyToBuy *ActiveCurrencyCode `xml:"CcyToBuy,omitempty"`

	// Account servicer is instructed to sell the indicated currency in order to obtain the necessary currency to fund the transaction within this instruction message.
	CurrencyToSell *ActiveCurrencyCode `xml:"CcyToSell,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification19 `xml:"SctyId,omitempty"`

	// Provides information about securities quantity linked to a corporate action option.
	SecuritiesQuantityOrInstructedAmount *SecuritiesQuantityOrAmount1Choice `xml:"SctiesQtyOrInstdAmt"`

	// Date/time at which the instructing party requests the instruction to be executed.
	ExecutionRequestedDateTime *DateAndDateTimeChoice `xml:"ExctnReqdDtTm,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate71 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice60 `xml:"PricDtls,omitempty"`

	// Provides additional information.
	AdditionalInformation *CorporateActionNarrative32 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionOption131) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption131) AddOptionType() *CorporateActionOption20Choice {
	c.OptionType = new(CorporateActionOption20Choice)
	return c.OptionType
}

func (c *CorporateActionOption131) AddFractionDisposition() *FractionDispositionType28Choice {
	c.FractionDisposition = new(FractionDispositionType28Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption131) AddChangeType() *CorporateActionChangeTypeFormat6Choice {
	newValue := new(CorporateActionChangeTypeFormat6Choice)
	c.ChangeType = append(c.ChangeType, newValue)
	return newValue
}

func (c *CorporateActionOption131) SetEligibleForCollateralIndicator(value string) {
	c.EligibleForCollateralIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption131) SetCurrencyToBuy(value string) {
	c.CurrencyToBuy = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption131) SetCurrencyToSell(value string) {
	c.CurrencyToSell = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption131) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption131) AddSecurityIdentification() *SecurityIdentification19 {
	c.SecurityIdentification = new(SecurityIdentification19)
	return c.SecurityIdentification
}

func (c *CorporateActionOption131) AddSecuritiesQuantityOrInstructedAmount() *SecuritiesQuantityOrAmount1Choice {
	c.SecuritiesQuantityOrInstructedAmount = new(SecuritiesQuantityOrAmount1Choice)
	return c.SecuritiesQuantityOrInstructedAmount
}

func (c *CorporateActionOption131) AddExecutionRequestedDateTime() *DateAndDateTimeChoice {
	c.ExecutionRequestedDateTime = new(DateAndDateTimeChoice)
	return c.ExecutionRequestedDateTime
}

func (c *CorporateActionOption131) AddRateAndAmountDetails() *CorporateActionRate71 {
	c.RateAndAmountDetails = new(CorporateActionRate71)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption131) AddPriceDetails() *CorporateActionPrice60 {
	c.PriceDetails = new(CorporateActionPrice60)
	return c.PriceDetails
}

func (c *CorporateActionOption131) AddAdditionalInformation() *CorporateActionNarrative32 {
	c.AdditionalInformation = new(CorporateActionNarrative32)
	return c.AdditionalInformation
}
