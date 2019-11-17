package model

// Amount of money for which goods or services are offered, sold, or bought.
type UnitPrice19 struct {

	// Type of price.
	PriceType *UnitPriceType2Choice `xml:"PricTp"`

	// Value of the price, that is, as a currency and value.
	Value *PriceValue1 `xml:"Val"`
}

func (u *UnitPrice19) AddPriceType() *UnitPriceType2Choice {
	u.PriceType = new(UnitPriceType2Choice)
	return u.PriceType
}

func (u *UnitPrice19) AddValue() *PriceValue1 {
	u.Value = new(PriceValue1)
	return u.Value
}
