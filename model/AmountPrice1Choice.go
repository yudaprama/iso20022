package model

// Choice between ways to express an amount.
type AmountPrice1Choice struct {

	// Amount.
	Amount *CurrencyAndAmount `xml:"Amt"`

	// Amount expressed as a unit price.
	UnitPrice *UnitPrice4 `xml:"UnitPric"`
}

func (a *AmountPrice1Choice) SetAmount(value, currency string) {
	a.Amount = NewCurrencyAndAmount(value, currency)
}

func (a *AmountPrice1Choice) AddUnitPrice() *UnitPrice4 {
	a.UnitPrice = new(UnitPrice4)
	return a.UnitPrice
}
