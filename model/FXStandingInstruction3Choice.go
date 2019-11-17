package model

// Choice of format for the FX Standing Instruction information.
type FXStandingInstruction3Choice struct {

	// Specifies whether the forex standing instruction in place should apply.
	Indicator *YesNoIndicator `xml:"Ind"`

	// FX Standing instruction information expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (f *FXStandingInstruction3Choice) SetIndicator(value string) {
	f.Indicator = (*YesNoIndicator)(&value)
}

func (f *FXStandingInstruction3Choice) AddProprietary() *GenericIdentification38 {
	f.Proprietary = new(GenericIdentification38)
	return f.Proprietary
}
