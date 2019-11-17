package model

// Provides information about the prices related to a corporate action option.
type PriceDetails11 struct {

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount or a number of points above an index, for example, reinvestment price, strike price and exercise price.
	GenericCashPricePaidPerProduct *PriceFormat6Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat7Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (p *PriceDetails11) AddGenericCashPricePaidPerProduct() *PriceFormat6Choice {
	p.GenericCashPricePaidPerProduct = new(PriceFormat6Choice)
	return p.GenericCashPricePaidPerProduct
}

func (p *PriceDetails11) AddGenericCashPriceReceivedPerProduct() *PriceFormat7Choice {
	p.GenericCashPriceReceivedPerProduct = new(PriceFormat7Choice)
	return p.GenericCashPriceReceivedPerProduct
}
