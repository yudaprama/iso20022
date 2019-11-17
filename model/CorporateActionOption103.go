package model

// Provides information about the corporate action option.
type CorporateActionOption103 struct {

	// Number identifying the available corporate action options.
	OptionNumber *OptionNumber1Choice `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption12Choice `xml:"OptnTp"`

	// Specifies how fractional amount/quantities are treated.
	FractionDisposition *FractionDispositionType17Choice `xml:"FrctnDspstn,omitempty"`

	// Type of changes affecting the security form.
	ChangeType []*CorporateActionChangeTypeFormat2Choice `xml:"ChngTp,omitempty"`

	// Specifies that the corporate action instruction is to be processed using the Available-for-Collateral pool.
	EligibleForCollateralIndicator *YesNoIndicator `xml:"ElgblForCollInd,omitempty"`

	// Account servicer is instructed to buy the indicated currency after the receipt of cash proceeds.
	CurrencyToBuy *ActiveCurrencyCode `xml:"CcyToBuy,omitempty"`

	// Account servicer is instructed to sell the indicated currency in order to obtain the necessary currency to fund the transaction within this instruction message.
	CurrencyToSell *ActiveCurrencyCode `xml:"CcyToSell,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification14 `xml:"SctyId,omitempty"`

	// Provides information about securities quantity linked to a corporate action option.
	SecuritiesQuantity *SecuritiesOption2 `xml:"SctiesQty"`

	// Date/time at which the instructing party requests the instruction to be executed.
	ExecutionRequestedDateTime *DateAndDateTimeChoice `xml:"ExctnReqdDtTm,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate47 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice44 `xml:"PricDtls,omitempty"`

	// Provides additional information.
	AdditionalInformation *CorporateActionNarrative8 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionOption103) AddOptionNumber() *OptionNumber1Choice {
	c.OptionNumber = new(OptionNumber1Choice)
	return c.OptionNumber
}

func (c *CorporateActionOption103) AddOptionType() *CorporateActionOption12Choice {
	c.OptionType = new(CorporateActionOption12Choice)
	return c.OptionType
}

func (c *CorporateActionOption103) AddFractionDisposition() *FractionDispositionType17Choice {
	c.FractionDisposition = new(FractionDispositionType17Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption103) AddChangeType() *CorporateActionChangeTypeFormat2Choice {
	newValue := new(CorporateActionChangeTypeFormat2Choice)
	c.ChangeType = append(c.ChangeType, newValue)
	return newValue
}

func (c *CorporateActionOption103) SetEligibleForCollateralIndicator(value string) {
	c.EligibleForCollateralIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption103) SetCurrencyToBuy(value string) {
	c.CurrencyToBuy = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption103) SetCurrencyToSell(value string) {
	c.CurrencyToSell = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption103) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption103) AddSecurityIdentification() *SecurityIdentification14 {
	c.SecurityIdentification = new(SecurityIdentification14)
	return c.SecurityIdentification
}

func (c *CorporateActionOption103) AddSecuritiesQuantity() *SecuritiesOption2 {
	c.SecuritiesQuantity = new(SecuritiesOption2)
	return c.SecuritiesQuantity
}

func (c *CorporateActionOption103) AddExecutionRequestedDateTime() *DateAndDateTimeChoice {
	c.ExecutionRequestedDateTime = new(DateAndDateTimeChoice)
	return c.ExecutionRequestedDateTime
}

func (c *CorporateActionOption103) AddRateAndAmountDetails() *CorporateActionRate47 {
	c.RateAndAmountDetails = new(CorporateActionRate47)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption103) AddPriceDetails() *CorporateActionPrice44 {
	c.PriceDetails = new(CorporateActionPrice44)
	return c.PriceDetails
}

func (c *CorporateActionOption103) AddAdditionalInformation() *CorporateActionNarrative8 {
	c.AdditionalInformation = new(CorporateActionNarrative8)
	return c.AdditionalInformation
}
