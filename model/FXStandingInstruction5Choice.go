package model

// Choice of format for the forex standing instruction information.
type FXStandingInstruction5Choice struct {

	// Specifies whether the forex standing instruction in place should apply.
	Indicator *YesNoIndicator `xml:"Ind"`

	// FX Standing instruction information expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (f *FXStandingInstruction5Choice) SetIndicator(value string) {
	f.Indicator = (*YesNoIndicator)(&value)
}

func (f *FXStandingInstruction5Choice) AddProprietary() *GenericIdentification47 {
	f.Proprietary = new(GenericIdentification47)
	return f.Proprietary
}
