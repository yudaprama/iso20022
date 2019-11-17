package model

// Information that describes a netting cut off held at a central system.
type CutOff1 struct {

	// Identification for the updated netting cut off.
	CutOffUpdateIdentification *Max35Text `xml:"CutOffUpdId"`

	// Currency linked to the netting cut off.
	Currency *ActiveCurrencyCode `xml:"Ccy"`

	// Cut off time value for the netting cut off.
	CutOffTime *ISOTime `xml:"CutOffTm"`

	// Specifies the offset in business days from the value date from which the netting cut off is to be applied.
	ValueDateOffset *DateOffsetText `xml:"ValDtOffset"`
}

func (c *CutOff1) SetCutOffUpdateIdentification(value string) {
	c.CutOffUpdateIdentification = (*Max35Text)(&value)
}

func (c *CutOff1) SetCurrency(value string) {
	c.Currency = (*ActiveCurrencyCode)(&value)
}

func (c *CutOff1) SetCutOffTime(value string) {
	c.CutOffTime = (*ISOTime)(&value)
}

func (c *CutOff1) SetValueDateOffset(value string) {
	c.ValueDateOffset = (*DateOffsetText)(&value)
}
