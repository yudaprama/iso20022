package model

// Choice of format for the forex cancellation information.
type FXCancellation3Choice struct {

	// Specifies whether the underlying forex transaction should also be cancelled. Yes means forex is to be cancelled. No means forex is to be retained.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Forex cancellation information expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (f *FXCancellation3Choice) SetIndicator(value string) {
	f.Indicator = (*YesNoIndicator)(&value)
}

func (f *FXCancellation3Choice) AddProprietary() *GenericIdentification30 {
	f.Proprietary = new(GenericIdentification30)
	return f.Proprietary
}
