package model

// Breakdown of the cash movements into a fund by order type, eg, order by quantity of units or amount of money.
type OriginalOrderQuantityType1 struct {

	// Order type, expressed as a code.
	Structured *OrderQuantityType1Code `xml:"Strd"`

	// Additional information about the order type.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (o *OriginalOrderQuantityType1) SetStructured(value string) {
	o.Structured = (*OrderQuantityType1Code)(&value)
}

func (o *OriginalOrderQuantityType1) SetAdditionalInformation(value string) {
	o.AdditionalInformation = (*Max350Text)(&value)
}
