package model

// Specifies the billing price of a service.
type BillingPrice1 struct {

	// Currency code in which the unit price and original charge price are expressed.
	Currency *ActiveOrHistoricCurrencyCode `xml:"Ccy,omitempty"`

	// Price per item or unit used to calculate the charge expressed in the pricing currency.
	UnitPrice *AmountAndDirection34 `xml:"UnitPric,omitempty"`

	// Identifies how the charge was calculated.
	//
	// Usage: The absence of this code assumes that the charge is calculated as the product of (volume x unit price).
	Method *BillingChargeMethod1Code `xml:"Mtd,omitempty"`

	// Indicates that the charge calculation is based on a particular rule. The rule name is carried here and is defined by the trading partners.
	Rule *Max20Text `xml:"Rule,omitempty"`
}

func (b *BillingPrice1) SetCurrency(value string) {
	b.Currency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (b *BillingPrice1) AddUnitPrice() *AmountAndDirection34 {
	b.UnitPrice = new(AmountAndDirection34)
	return b.UnitPrice
}

func (b *BillingPrice1) SetMethod(value string) {
	b.Method = (*BillingChargeMethod1Code)(&value)
}

func (b *BillingPrice1) SetRule(value string) {
	b.Rule = (*Max20Text)(&value)
}
