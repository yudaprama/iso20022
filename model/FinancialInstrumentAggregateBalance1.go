package model

// Aggregated position of holdings held in a securities account for a specified financial instrument.
type FinancialInstrumentAggregateBalance1 struct {

	// Date of the line of holding in the statement.
	ItemDate *ISODate `xml:"ItmDt"`

	// Balances and sub-balances attributed to the holding.
	Holdings *FinancialInstrumentAggregateBalance1Choice `xml:"Hldgs"`

	// Details on the price value, type and source.
	Price []*Price6 `xml:"Pric,omitempty"`
}

func (f *FinancialInstrumentAggregateBalance1) SetItemDate(value string) {
	f.ItemDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAggregateBalance1) AddHoldings() *FinancialInstrumentAggregateBalance1Choice {
	f.Holdings = new(FinancialInstrumentAggregateBalance1Choice)
	return f.Holdings
}

func (f *FinancialInstrumentAggregateBalance1) AddPrice() *Price6 {
	newValue := new(Price6)
	f.Price = append(f.Price, newValue)
	return newValue
}
