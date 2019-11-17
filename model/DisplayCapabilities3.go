package model

// Capabilities of the display components performing the transaction.
type DisplayCapabilities3 struct {

	// Destination of the message to present.
	Destination *UserInterface1Code `xml:"Dstn"`

	// Available message format.
	AvailableFormat []*OutputFormat1Code `xml:"AvlblFrmt"`

	// Number of lines of the display.
	NumberOfLines *Number `xml:"NbOfLines,omitempty"`

	// Number of columns of the display or printer.
	LineWidth *Number `xml:"LineWidth,omitempty"`

	// Available language for the message. Reference ISO 639-1 (alpha-2) et ISO 639-2 (alpha-3).
	AvailableLanguage []*LanguageCode `xml:"AvlblLang,omitempty"`
}

func (d *DisplayCapabilities3) SetDestination(value string) {
	d.Destination = (*UserInterface1Code)(&value)
}

func (d *DisplayCapabilities3) AddAvailableFormat(value string) {
	d.AvailableFormat = append(d.AvailableFormat, (*OutputFormat1Code)(&value))
}

func (d *DisplayCapabilities3) SetNumberOfLines(value string) {
	d.NumberOfLines = (*Number)(&value)
}

func (d *DisplayCapabilities3) SetLineWidth(value string) {
	d.LineWidth = (*Number)(&value)
}

func (d *DisplayCapabilities3) AddAvailableLanguage(value string) {
	d.AvailableLanguage = append(d.AvailableLanguage, (*LanguageCode)(&value))
}
