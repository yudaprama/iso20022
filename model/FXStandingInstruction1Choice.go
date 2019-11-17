package model

// Choice of format for the FX Standing Instruction information.
type FXStandingInstruction1Choice struct {

	// Specifies whether the forex standing instruction in place should apply.
	Indicator *YesNoIndicator `xml:"Ind"`

	// FX Standing instruction information expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (f *FXStandingInstruction1Choice) SetIndicator(value string) {
	f.Indicator = (*YesNoIndicator)(&value)
}

func (f *FXStandingInstruction1Choice) AddProprietary() *GenericIdentification20 {
	f.Proprietary = new(GenericIdentification20)
	return f.Proprietary
}
