package model

// Choice of format for the tracking information.
type Tracking3Choice struct {

	// Specifies whether the loan and/or collateral is tracked.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Tracking information expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (t *Tracking3Choice) SetIndicator(value string) {
	t.Indicator = (*YesNoIndicator)(&value)
}

func (t *Tracking3Choice) AddProprietary() *GenericIdentification38 {
	t.Proprietary = new(GenericIdentification38)
	return t.Proprietary
}
