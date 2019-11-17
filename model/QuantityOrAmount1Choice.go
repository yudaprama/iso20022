package model

// Choice between a quantity expressed in units, face amount or amortised amount and a quantity expressed as an amount.
type QuantityOrAmount1Choice struct {

	// Quantity is expressed in units, face amount or amortised amount.
	Quantity *FinancialInstrumentQuantityChoice `xml:"Qty"`

	// Quantity is expressed as an amount.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`
}

func (q *QuantityOrAmount1Choice) AddQuantity() *FinancialInstrumentQuantityChoice {
	q.Quantity = new(FinancialInstrumentQuantityChoice)
	return q.Quantity
}

func (q *QuantityOrAmount1Choice) SetAmount(value, currency string) {
	q.Amount = NewActiveCurrencyAndAmount(value, currency)
}
