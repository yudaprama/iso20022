package model

// Choice of format for the tracking information.
type Tracking4Choice struct {

	// Specifies whether the loan and/or collateral is tracked.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Tracking information expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (t *Tracking4Choice) SetIndicator(value string) {
	t.Indicator = (*YesNoIndicator)(&value)
}

func (t *Tracking4Choice) AddProprietary() *GenericIdentification30 {
	t.Proprietary = new(GenericIdentification30)
	return t.Proprietary
}
