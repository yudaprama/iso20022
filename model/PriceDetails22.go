package model

// Provides information about the prices related to a corporate action option.
type PriceDetails22 struct {

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount or a number of points above an index, for example, reinvestment price, strike price and exercise price.
	GenericCashPricePaidPerProduct *PriceFormat44Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat47Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (p *PriceDetails22) AddGenericCashPricePaidPerProduct() *PriceFormat44Choice {
	p.GenericCashPricePaidPerProduct = new(PriceFormat44Choice)
	return p.GenericCashPricePaidPerProduct
}

func (p *PriceDetails22) AddGenericCashPriceReceivedPerProduct() *PriceFormat47Choice {
	p.GenericCashPriceReceivedPerProduct = new(PriceFormat47Choice)
	return p.GenericCashPriceReceivedPerProduct
}
