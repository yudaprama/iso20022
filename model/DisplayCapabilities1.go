package model

// The capabilities of the display components performing the transaction.
type DisplayCapabilities1 struct {

	// Type of display (for example merchant or cardholder).
	DisplayType *UserInterface2Code `xml:"DispTp"`

	// Number of lines of the display component.
	NumberOfLines *Max3NumericText `xml:"NbOfLines"`

	// Number of columns of the display component.
	LineWidth *Max3NumericText `xml:"LineWidth"`
}

func (d *DisplayCapabilities1) SetDisplayType(value string) {
	d.DisplayType = (*UserInterface2Code)(&value)
}

func (d *DisplayCapabilities1) SetNumberOfLines(value string) {
	d.NumberOfLines = (*Max3NumericText)(&value)
}

func (d *DisplayCapabilities1) SetLineWidth(value string) {
	d.LineWidth = (*Max3NumericText)(&value)
}
