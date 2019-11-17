package model

// Choice of value type.
type YieldedOrValueType1Choice struct {

	// Indicates whether the price is expressed as a yield.
	Yielded *YesNoIndicator `xml:"Yldd"`

	// Type of value in which the price is expressed.
	ValueType *PriceValueType1Code `xml:"ValTp"`
}

func (y *YieldedOrValueType1Choice) SetYielded(value string) {
	y.Yielded = (*YesNoIndicator)(&value)
}

func (y *YieldedOrValueType1Choice) SetValueType(value string) {
	y.ValueType = (*PriceValueType1Code)(&value)
}
