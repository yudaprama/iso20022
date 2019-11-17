package model

// Choice of price type.
type PriceType1Choice struct {

	// Last reported price of a financial instrument in a market, determined by supply and demand.
	Market *Price2 `xml:"Mkt"`

	// Estimated price, for valuation purposes.
	Indicative *Price2 `xml:"Indctv"`
}

func (p *PriceType1Choice) AddMarket() *Price2 {
	p.Market = new(Price2)
	return p.Market
}

func (p *PriceType1Choice) AddIndicative() *Price2 {
	p.Indicative = new(Price2)
	return p.Indicative
}
