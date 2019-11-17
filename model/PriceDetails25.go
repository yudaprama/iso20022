package model

// Provides information about the prices related to a corporate action option.
type PriceDetails25 struct {

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount or a number of points above an index, for example, reinvestment price, strike price and exercise price.
	GenericCashPricePaidPerProduct *PriceFormat59Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat60Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (p *PriceDetails25) AddGenericCashPricePaidPerProduct() *PriceFormat59Choice {
	p.GenericCashPricePaidPerProduct = new(PriceFormat59Choice)
	return p.GenericCashPricePaidPerProduct
}

func (p *PriceDetails25) AddGenericCashPriceReceivedPerProduct() *PriceFormat60Choice {
	p.GenericCashPriceReceivedPerProduct = new(PriceFormat60Choice)
	return p.GenericCashPriceReceivedPerProduct
}
