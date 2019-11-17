package model

// Choice of format for the tracking information.
type Tracking5Choice struct {

	// Specifies whether the loan and/or collateral is tracked.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Tracking information expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (t *Tracking5Choice) SetIndicator(value string) {
	t.Indicator = (*YesNoIndicator)(&value)
}

func (t *Tracking5Choice) AddProprietary() *GenericIdentification47 {
	t.Proprietary = new(GenericIdentification47)
	return t.Proprietary
}
