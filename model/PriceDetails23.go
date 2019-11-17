package model

// Provides information about the prices related to a corporate action option.
type PriceDetails23 struct {

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount or a number of points above an index, for example, reinvestment price, strike price and exercise price.
	GenericCashPricePaidPerProduct *PriceFormat51Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat48Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (p *PriceDetails23) AddGenericCashPricePaidPerProduct() *PriceFormat51Choice {
	p.GenericCashPricePaidPerProduct = new(PriceFormat51Choice)
	return p.GenericCashPricePaidPerProduct
}

func (p *PriceDetails23) AddGenericCashPriceReceivedPerProduct() *PriceFormat48Choice {
	p.GenericCashPriceReceivedPerProduct = new(PriceFormat48Choice)
	return p.GenericCashPriceReceivedPerProduct
}
