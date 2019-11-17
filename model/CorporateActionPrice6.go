package model

// Specifies prices related to a corporate action option.
type CorporateActionPrice6 struct {

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat11Choice `xml:"CshInLieuOfShrPric,omitempty"`

	// 1. Price at which security will be purchased/sold if warrant is exercised, either as an actual amount or a percentage.
	// 2. Price at which a bond is converted to underlying security either as an actual amount or a percentage.
	// 3. Strike price of an option, represented either as an actual amount, a percentage or a number of points above an index.
	ExercisePrice *PriceFormat8Choice `xml:"ExrcPric,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct []*PriceFormat9Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`

	// Amount of money required per over-subscribed equity as defined by the issuer.
	OverSubscriptionDepositPrice *PriceFormat11Choice `xml:"OverSbcptDpstPric,omitempty"`

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount, for example, reinvestment price.
	GenericCashPricePaidPerProduct *PriceFormat11Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Amount included in the dividend/NAV that is identified as gains directly or indirectly derived from interest payments within the scope of the EU Savings directive.
	TaxableIncomePerDividendShare *AmountPrice3 `xml:"TaxblIncmPerDvddShr,omitempty"`
}

func (c *CorporateActionPrice6) AddCashInLieuOfSharePrice() *PriceFormat11Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat11Choice)
	return c.CashInLieuOfSharePrice
}

func (c *CorporateActionPrice6) AddExercisePrice() *PriceFormat8Choice {
	c.ExercisePrice = new(PriceFormat8Choice)
	return c.ExercisePrice
}

func (c *CorporateActionPrice6) AddGenericCashPriceReceivedPerProduct() *PriceFormat9Choice {
	newValue := new(PriceFormat9Choice)
	c.GenericCashPriceReceivedPerProduct = append(c.GenericCashPriceReceivedPerProduct, newValue)
	return newValue
}

func (c *CorporateActionPrice6) AddOverSubscriptionDepositPrice() *PriceFormat11Choice {
	c.OverSubscriptionDepositPrice = new(PriceFormat11Choice)
	return c.OverSubscriptionDepositPrice
}

func (c *CorporateActionPrice6) AddGenericCashPricePaidPerProduct() *PriceFormat11Choice {
	c.GenericCashPricePaidPerProduct = new(PriceFormat11Choice)
	return c.GenericCashPricePaidPerProduct
}

func (c *CorporateActionPrice6) AddTaxableIncomePerDividendShare() *AmountPrice3 {
	c.TaxableIncomePerDividendShare = new(AmountPrice3)
	return c.TaxableIncomePerDividendShare
}
