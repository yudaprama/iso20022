package model

// Choice between a type of discount or a type of charge.
type DiscountOrChargeType1Choice struct {

	// Type of charge.
	ChargeType *ChargeTypeFormat3Choice `xml:"ChrgTp"`

	// Type of discount.
	DiscountType *DiscountTypeFormat1Choice `xml:"DscntTp"`
}

func (d *DiscountOrChargeType1Choice) AddChargeType() *ChargeTypeFormat3Choice {
	d.ChargeType = new(ChargeTypeFormat3Choice)
	return d.ChargeType
}

func (d *DiscountOrChargeType1Choice) AddDiscountType() *DiscountTypeFormat1Choice {
	d.DiscountType = new(DiscountTypeFormat1Choice)
	return d.DiscountType
}
