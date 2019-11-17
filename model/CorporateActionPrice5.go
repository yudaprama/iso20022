package model

// Specifies prices.
type CorporateActionPrice5 struct {

	// 1. Price at which security will be purchased/sold if warrant is exercised, either as an actual amount or a percentage.
	// 2. Price at which a bond is converted to underlying security either as an actual amount or a percentage.
	// 3. Strike price of an option, represented either as an actual amount, a percentage or a number of points above an index.
	ExercisePrice *PriceFormat6Choice `xml:"ExrcPric,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct []*PriceFormat7Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount, for example, reinvestment price.
	GenericCashPricePaidPerProduct *PriceFormat5Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Amount included in the dividend/NAV that is identified as gains directly or indirectly derived from interest payments within the scope of the EU Savings directive.
	TaxableIncomePerDividendShare *AmountPrice3 `xml:"TaxblIncmPerDvddShr,omitempty"`

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat5Choice `xml:"CshInLieuOfShrPric,omitempty"`

	// Amount of money required per over-subscribed equity as defined by the issuer.
	OverSubscriptionDepositPrice *PriceFormat5Choice `xml:"OverSbcptDpstPric,omitempty"`
}

func (c *CorporateActionPrice5) AddExercisePrice() *PriceFormat6Choice {
	c.ExercisePrice = new(PriceFormat6Choice)
	return c.ExercisePrice
}

func (c *CorporateActionPrice5) AddGenericCashPriceReceivedPerProduct() *PriceFormat7Choice {
	newValue := new(PriceFormat7Choice)
	c.GenericCashPriceReceivedPerProduct = append(c.GenericCashPriceReceivedPerProduct, newValue)
	return newValue
}

func (c *CorporateActionPrice5) AddGenericCashPricePaidPerProduct() *PriceFormat5Choice {
	c.GenericCashPricePaidPerProduct = new(PriceFormat5Choice)
	return c.GenericCashPricePaidPerProduct
}

func (c *CorporateActionPrice5) AddTaxableIncomePerDividendShare() *AmountPrice3 {
	c.TaxableIncomePerDividendShare = new(AmountPrice3)
	return c.TaxableIncomePerDividendShare
}

func (c *CorporateActionPrice5) AddCashInLieuOfSharePrice() *PriceFormat5Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat5Choice)
	return c.CashInLieuOfSharePrice
}

func (c *CorporateActionPrice5) AddOverSubscriptionDepositPrice() *PriceFormat5Choice {
	c.OverSubscriptionDepositPrice = new(PriceFormat5Choice)
	return c.OverSubscriptionDepositPrice
}
