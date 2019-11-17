package model

// Provides information about the prices related to a corporate action option.
type PriceDetails24 struct {

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount or a number of points above an index, for example, reinvestment price, strike price and exercise price.
	GenericCashPricePaidPerProduct *PriceFormat55Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat56Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (p *PriceDetails24) AddGenericCashPricePaidPerProduct() *PriceFormat55Choice {
	p.GenericCashPricePaidPerProduct = new(PriceFormat55Choice)
	return p.GenericCashPricePaidPerProduct
}

func (p *PriceDetails24) AddGenericCashPriceReceivedPerProduct() *PriceFormat56Choice {
	p.GenericCashPriceReceivedPerProduct = new(PriceFormat56Choice)
	return p.GenericCashPriceReceivedPerProduct
}
