package model

// Provides information about the prices related to a corporate action option.
type PriceDetails7 struct {

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount, for example, reinvestment price.
	GenericCashPricePaidPerProduct *PriceFormat6Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat7Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (p *PriceDetails7) AddGenericCashPricePaidPerProduct() *PriceFormat6Choice {
	p.GenericCashPricePaidPerProduct = new(PriceFormat6Choice)
	return p.GenericCashPricePaidPerProduct
}

func (p *PriceDetails7) AddGenericCashPriceReceivedPerProduct() *PriceFormat7Choice {
	p.GenericCashPriceReceivedPerProduct = new(PriceFormat7Choice)
	return p.GenericCashPriceReceivedPerProduct
}
