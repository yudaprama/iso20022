package model

// Product to purchase.
type Product5 struct {

	// Product code of the item.
	ProductCode *Max70Text `xml:"PdctCd"`

	// Additional product code related to the product.
	AdditionalProductCode *Max70Text `xml:"AddtlPdctCd,omitempty"`

	// Amount limit for the product.
	AmountLimit *ImpliedCurrencyAndAmount `xml:"AmtLmt,omitempty"`

	// Quantity limit for the product.
	QuantityLimit *DecimalNumber `xml:"QtyLmt,omitempty"`

	// Unit of measure of the item purchased.
	UnitOfMeasure *UnitOfMeasure6Code `xml:"UnitOfMeasr,omitempty"`
}

func (p *Product5) SetProductCode(value string) {
	p.ProductCode = (*Max70Text)(&value)
}

func (p *Product5) SetAdditionalProductCode(value string) {
	p.AdditionalProductCode = (*Max70Text)(&value)
}

func (p *Product5) SetAmountLimit(value, currency string) {
	p.AmountLimit = NewImpliedCurrencyAndAmount(value, currency)
}

func (p *Product5) SetQuantityLimit(value string) {
	p.QuantityLimit = (*DecimalNumber)(&value)
}

func (p *Product5) SetUnitOfMeasure(value string) {
	p.UnitOfMeasure = (*UnitOfMeasure6Code)(&value)
}
