package model

// Choice between a standard code or a proprietary code for specifying the type of discount.
type DiscountTypeFormat1Choice struct {

	// Standard code to specify the type of discount.
	Code *DiscountType1Code `xml:"Cd"`

	// Proprietary code for specifying the type of discount.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (d *DiscountTypeFormat1Choice) SetCode(value string) {
	d.Code = (*DiscountType1Code)(&value)
}

func (d *DiscountTypeFormat1Choice) AddProprietary() *GenericIdentification13 {
	d.Proprietary = new(GenericIdentification13)
	return d.Proprietary
}
