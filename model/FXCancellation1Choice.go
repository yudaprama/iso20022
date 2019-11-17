package model

// Choice of format for the FX cancellation information.
type FXCancellation1Choice struct {

	// Specifies whether the underlying FX transaction should also be cancelled. Yes means FX is to be cancelled. No means FX is to be retained.
	Indicator *YesNoIndicator `xml:"Ind"`

	// FX Cancellation information expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (f *FXCancellation1Choice) SetIndicator(value string) {
	f.Indicator = (*YesNoIndicator)(&value)
}

func (f *FXCancellation1Choice) AddProprietary() *GenericIdentification20 {
	f.Proprietary = new(GenericIdentification20)
	return f.Proprietary
}
