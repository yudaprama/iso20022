package model

// Specifies prices.
type CorporateActionPrice1 struct {

	// 1. Price at which security will be purchased/sold if warrant is exercised, either as an actual amount or a percentage.
	// 2. Price at which a bond is converted to underlying security either as an actual amount or a percentage.
	// 3. Strike price of an option, represented either as an actual amount, a percentage or or a number of points above an index.
	ExercisePrice *PriceFormat4Choice `xml:"ExrcPric,omitempty"`

	// Initial issue price of a financial instrument.
	IssuePrice *PriceFormat2Choice `xml:"IssePric,omitempty"`

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat2Choice `xml:"CshInLieuOfShrPric,omitempty"`

	// Amount included in the dividend/NAV that is identified as gains directly or indirectly derived from interest payments within the scope of the EU Savings directive.
	TaxableIncomePerDividendShare *AmountPrice1 `xml:"TaxblIncmPerDvddShr,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, eg, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat1Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount, eg, reinvestment price.
	GenericCashPricePaidPerProduct *PriceFormat2Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Amount of money required per over-subscribed equity as defined by the issuer.
	OverSubscriptionDepositPrice *PriceFormat2Choice `xml:"OverSbcptDpstPric,omitempty"`
}

func (c *CorporateActionPrice1) AddExercisePrice() *PriceFormat4Choice {
	c.ExercisePrice = new(PriceFormat4Choice)
	return c.ExercisePrice
}

func (c *CorporateActionPrice1) AddIssuePrice() *PriceFormat2Choice {
	c.IssuePrice = new(PriceFormat2Choice)
	return c.IssuePrice
}

func (c *CorporateActionPrice1) AddCashInLieuOfSharePrice() *PriceFormat2Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat2Choice)
	return c.CashInLieuOfSharePrice
}

func (c *CorporateActionPrice1) AddTaxableIncomePerDividendShare() *AmountPrice1 {
	c.TaxableIncomePerDividendShare = new(AmountPrice1)
	return c.TaxableIncomePerDividendShare
}

func (c *CorporateActionPrice1) AddGenericCashPriceReceivedPerProduct() *PriceFormat1Choice {
	c.GenericCashPriceReceivedPerProduct = new(PriceFormat1Choice)
	return c.GenericCashPriceReceivedPerProduct
}

func (c *CorporateActionPrice1) AddGenericCashPricePaidPerProduct() *PriceFormat2Choice {
	c.GenericCashPricePaidPerProduct = new(PriceFormat2Choice)
	return c.GenericCashPricePaidPerProduct
}

func (c *CorporateActionPrice1) AddOverSubscriptionDepositPrice() *PriceFormat2Choice {
	c.OverSubscriptionDepositPrice = new(PriceFormat2Choice)
	return c.OverSubscriptionDepositPrice
}
