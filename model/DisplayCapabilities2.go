package model

// The capabilities of the display components performing the transaction.
type DisplayCapabilities2 struct {

	// Type of display (for example merchant or cardholder).
	DisplayType *UserInterface2Code `xml:"DispTp"`

	// Number of lines of the display component.
	NumberOfLines *Number `xml:"NbOfLines"`

	// Number of columns of the display component.
	LineWidth *Number `xml:"LineWidth"`
}

func (d *DisplayCapabilities2) SetDisplayType(value string) {
	d.DisplayType = (*UserInterface2Code)(&value)
}

func (d *DisplayCapabilities2) SetNumberOfLines(value string) {
	d.NumberOfLines = (*Number)(&value)
}

func (d *DisplayCapabilities2) SetLineWidth(value string) {
	d.LineWidth = (*Number)(&value)
}
