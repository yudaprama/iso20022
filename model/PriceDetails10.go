package model

// Provides information about the prices related to a corporate action option.
type PriceDetails10 struct {

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount or a number of points above an index, for example, reinvestment price, strike price and exercise price.
	GenericCashPricePaidPerProduct *PriceFormat23Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat20Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (p *PriceDetails10) AddGenericCashPricePaidPerProduct() *PriceFormat23Choice {
	p.GenericCashPricePaidPerProduct = new(PriceFormat23Choice)
	return p.GenericCashPricePaidPerProduct
}

func (p *PriceDetails10) AddGenericCashPriceReceivedPerProduct() *PriceFormat20Choice {
	p.GenericCashPriceReceivedPerProduct = new(PriceFormat20Choice)
	return p.GenericCashPriceReceivedPerProduct
}
