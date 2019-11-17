package model

// Net position of a segregated holding, in a single security, within the overall position held in a securities account. A securities balance is calculated from the sum of securities' receipts minus the sum of securities' deliveries.
type Balance8 struct {

	// Indication that the position is short or long.
	ShortLongIndicator *ShortLong1Code `xml:"ShrtLngInd,omitempty"`

	// Total quantity of financial instruments of the balance.
	Quantity *BalanceQuantity8Choice `xml:"Qty"`
}

func (b *Balance8) SetShortLongIndicator(value string) {
	b.ShortLongIndicator = (*ShortLong1Code)(&value)
}

func (b *Balance8) AddQuantity() *BalanceQuantity8Choice {
	b.Quantity = new(BalanceQuantity8Choice)
	return b.Quantity
}
