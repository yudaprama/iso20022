package model

// Provides information about the prices related to a corporate action option.
type PriceDetails2 struct {

	// 1. Price at which security will be purchased/sold if warrant is exercised, either as an actual amount or a percentage.
	// 2. Price at which a bond is converted to underlying security either as an actual amount or a percentage.
	// 3. Strike price of an option, represented either as an actual amount, a percentage or a number of points above an index.
	ExercisePrice *PriceFormat6Choice `xml:"ExrcPric,omitempty"`

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount, for example, reinvestment price.
	GenericCashPricePaidPerProduct *PriceFormat5Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat7Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (p *PriceDetails2) AddExercisePrice() *PriceFormat6Choice {
	p.ExercisePrice = new(PriceFormat6Choice)
	return p.ExercisePrice
}

func (p *PriceDetails2) AddGenericCashPricePaidPerProduct() *PriceFormat5Choice {
	p.GenericCashPricePaidPerProduct = new(PriceFormat5Choice)
	return p.GenericCashPricePaidPerProduct
}

func (p *PriceDetails2) AddGenericCashPriceReceivedPerProduct() *PriceFormat7Choice {
	p.GenericCashPriceReceivedPerProduct = new(PriceFormat7Choice)
	return p.GenericCashPriceReceivedPerProduct
}
