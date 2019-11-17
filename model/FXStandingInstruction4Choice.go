package model

// Choice of format for the forex standing instruction information.
type FXStandingInstruction4Choice struct {

	// Specifies whether the forex standing instruction in place should apply.
	Indicator *YesNoIndicator `xml:"Ind"`

	// FX Standing instruction information expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (f *FXStandingInstruction4Choice) SetIndicator(value string) {
	f.Indicator = (*YesNoIndicator)(&value)
}

func (f *FXStandingInstruction4Choice) AddProprietary() *GenericIdentification30 {
	f.Proprietary = new(GenericIdentification30)
	return f.Proprietary
}
