package model

// Specifies the amount type.
type DiscountAmountType1Choice struct {

	// Specifies the amount type, in a coded form.
	Code *ExternalDiscountAmountType1Code `xml:"Cd"`

	// Specifies the amount type, in a free-text form.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (d *DiscountAmountType1Choice) SetCode(value string) {
	d.Code = (*ExternalDiscountAmountType1Code)(&value)
}

func (d *DiscountAmountType1Choice) SetProprietary(value string) {
	d.Proprietary = (*Max35Text)(&value)
}
