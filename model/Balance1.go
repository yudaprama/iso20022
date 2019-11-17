package model

// Net position of a segregated holding, in a single security, within the overall position held in a securities account. A securities balance is calculated from the sum of securities' receipts minus the sum of securities' deliveries.
type Balance1 struct {

	// Indication that the position is short or long.
	ShortLongIndicator *ShortLong1Code `xml:"ShrtLngInd"`

	// Total quantity of financial instruments of the balance.
	Quantity *BalanceQuantity4Choice `xml:"Qty"`
}

func (b *Balance1) SetShortLongIndicator(value string) {
	b.ShortLongIndicator = (*ShortLong1Code)(&value)
}

func (b *Balance1) AddQuantity() *BalanceQuantity4Choice {
	b.Quantity = new(BalanceQuantity4Choice)
	return b.Quantity
}
