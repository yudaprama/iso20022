package model

// Choice of format for the tracking information.
type Tracking1Choice struct {

	// Specifies whether the loan and/or collateral is tracked.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Tracking information expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (t *Tracking1Choice) SetIndicator(value string) {
	t.Indicator = (*YesNoIndicator)(&value)
}

func (t *Tracking1Choice) AddProprietary() *GenericIdentification20 {
	t.Proprietary = new(GenericIdentification20)
	return t.Proprietary
}
