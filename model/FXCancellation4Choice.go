package model

// Choice of format for the forex cancellation information.
type FXCancellation4Choice struct {

	// Specifies whether the underlying forex transaction should also be cancelled. Yes means forex is to be cancelled. No means forex is to be retained.
	Indicator *YesNoIndicator `xml:"Ind"`

	// Forex cancellation information expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (f *FXCancellation4Choice) SetIndicator(value string) {
	f.Indicator = (*YesNoIndicator)(&value)
}

func (f *FXCancellation4Choice) AddProprietary() *GenericIdentification47 {
	f.Proprietary = new(GenericIdentification47)
	return f.Proprietary
}
