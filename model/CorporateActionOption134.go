package model

// Provides information about the corporate action option.
type CorporateActionOption134 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption29Choice `xml:"OptnTp"`

	// Specifies how fractional amount/quantities are treated.
	FractionDisposition *FractionDispositionType29Choice `xml:"FrctnDspstn,omitempty"`

	// Type of changes affecting the security form.
	ChangeType []*CorporateActionChangeTypeFormat7Choice `xml:"ChngTp,omitempty"`

	// Specifies that the corporate action instruction is to be processed using the Available-for-Collateral pool.
	EligibleForCollateralIndicator *YesNoIndicator `xml:"ElgblForCollInd,omitempty"`

	// Account servicer is instructed to buy the indicated currency after the receipt of cash proceeds.
	CurrencyToBuy *ActiveCurrencyCode `xml:"CcyToBuy,omitempty"`

	// Account servicer is instructed to sell the indicated currency in order to obtain the necessary currency to fund the transaction within this instruction message.
	CurrencyToSell *ActiveCurrencyCode `xml:"CcyToSell,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification20 `xml:"SctyId,omitempty"`

	// Provides information about securities quantity linked to a corporate action option.
	SecuritiesQuantityOrInstructedAmount *SecuritiesQuantityOrAmount1Choice `xml:"SctiesQtyOrInstdAmt"`

	// Date/time at which the instructing party requests the instruction to be executed.
	ExecutionRequestedDateTime *DateAndDateTimeChoice `xml:"ExctnReqdDtTm,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate73 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice62 `xml:"PricDtls,omitempty"`

	// Provides additional information.
	AdditionalInformation *CorporateActionNarrative33 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionOption134) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption134) AddOptionType() *CorporateActionOption29Choice {
	c.OptionType = new(CorporateActionOption29Choice)
	return c.OptionType
}

func (c *CorporateActionOption134) AddFractionDisposition() *FractionDispositionType29Choice {
	c.FractionDisposition = new(FractionDispositionType29Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption134) AddChangeType() *CorporateActionChangeTypeFormat7Choice {
	newValue := new(CorporateActionChangeTypeFormat7Choice)
	c.ChangeType = append(c.ChangeType, newValue)
	return newValue
}

func (c *CorporateActionOption134) SetEligibleForCollateralIndicator(value string) {
	c.EligibleForCollateralIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption134) SetCurrencyToBuy(value string) {
	c.CurrencyToBuy = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption134) SetCurrencyToSell(value string) {
	c.CurrencyToSell = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption134) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption134) AddSecurityIdentification() *SecurityIdentification20 {
	c.SecurityIdentification = new(SecurityIdentification20)
	return c.SecurityIdentification
}

func (c *CorporateActionOption134) AddSecuritiesQuantityOrInstructedAmount() *SecuritiesQuantityOrAmount1Choice {
	c.SecuritiesQuantityOrInstructedAmount = new(SecuritiesQuantityOrAmount1Choice)
	return c.SecuritiesQuantityOrInstructedAmount
}

func (c *CorporateActionOption134) AddExecutionRequestedDateTime() *DateAndDateTimeChoice {
	c.ExecutionRequestedDateTime = new(DateAndDateTimeChoice)
	return c.ExecutionRequestedDateTime
}

func (c *CorporateActionOption134) AddRateAndAmountDetails() *CorporateActionRate73 {
	c.RateAndAmountDetails = new(CorporateActionRate73)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption134) AddPriceDetails() *CorporateActionPrice62 {
	c.PriceDetails = new(CorporateActionPrice62)
	return c.PriceDetails
}

func (c *CorporateActionOption134) AddAdditionalInformation() *CorporateActionNarrative33 {
	c.AdditionalInformation = new(CorporateActionNarrative33)
	return c.AdditionalInformation
}
