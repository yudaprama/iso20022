package model

// Amount of money for which goods or services are offered, sold, or bought.
type UnitPrice11 struct {

	// Type and information about a price.
	Type *TypeOfPrice10Code `xml:"Tp"`

	// Type and information about a price.
	ExtendedType *Extended350Code `xml:"XtndedTp"`

	// Value of the price, eg, as a currency and value.
	Value *PriceValue1 `xml:"Val"`

	// Type of pricing calculation method.
	PriceMethod *PriceMethod1Code `xml:"PricMtd,omitempty"`
}

func (u *UnitPrice11) SetType(value string) {
	u.Type = (*TypeOfPrice10Code)(&value)
}

func (u *UnitPrice11) SetExtendedType(value string) {
	u.ExtendedType = (*Extended350Code)(&value)
}

func (u *UnitPrice11) AddValue() *PriceValue1 {
	u.Value = new(PriceValue1)
	return u.Value
}

func (u *UnitPrice11) SetPriceMethod(value string) {
	u.PriceMethod = (*PriceMethod1Code)(&value)
}
