package model

// Provides information about the prices related to a corporate action option.
type PriceDetails3 struct {

	// 1. Price at which security will be purchased/sold if warrant is exercised, either as an actual amount or a percentage.
	// 2. Price at which a bond is converted to underlying security either as an actual amount or a percentage.
	// 3. Strike price of an option, represented either as an actual amount, a percentage or a number of points above an index.
	ExercisePrice *PriceFormat23Choice `xml:"ExrcPric,omitempty"`

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount, for example, reinvestment price.
	GenericCashPricePaidPerProduct *PriceFormat19Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat22Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (p *PriceDetails3) AddExercisePrice() *PriceFormat23Choice {
	p.ExercisePrice = new(PriceFormat23Choice)
	return p.ExercisePrice
}

func (p *PriceDetails3) AddGenericCashPricePaidPerProduct() *PriceFormat19Choice {
	p.GenericCashPricePaidPerProduct = new(PriceFormat19Choice)
	return p.GenericCashPricePaidPerProduct
}

func (p *PriceDetails3) AddGenericCashPriceReceivedPerProduct() *PriceFormat22Choice {
	p.GenericCashPriceReceivedPerProduct = new(PriceFormat22Choice)
	return p.GenericCashPriceReceivedPerProduct
}
