package model

// Choice between an number and a option number in a code format.
type OptionNumber1Choice struct {

	// Number identifying the available corporate action options.
	Number *Exact3NumericText `xml:"Nb"`

	// Code identifying special corporate action option numbers.
	Code *OptionNumber1Code `xml:"Cd"`
}

func (o *OptionNumber1Choice) SetNumber(value string) {
	o.Number = (*Exact3NumericText)(&value)
}

func (o *OptionNumber1Choice) SetCode(value string) {
	o.Code = (*OptionNumber1Code)(&value)
}
