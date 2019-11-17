package model

// Choice of price type.
type PriceType2Choice struct {

	// Last reported price of a financial instrument in a market, determined by supply and demand.
	Market *Price3 `xml:"Mkt"`

	// Estimated price, for valuation purposes.
	Indicative *Price3 `xml:"Indctv"`
}

func (p *PriceType2Choice) AddMarket() *Price3 {
	p.Market = new(Price3)
	return p.Market
}

func (p *PriceType2Choice) AddIndicative() *Price3 {
	p.Indicative = new(Price3)
	return p.Indicative
}
