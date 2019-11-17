package model

// Additional information that can not be captured in the structured fields and/or any other specific block.
type Extension1 struct {

	// Name qualifying the information provided in the Text field, and place where this information should be inserted.
	PlaceAndName *Max350Text `xml:"PlcAndNm"`

	// Text of the extension.
	Text *Max350Text `xml:"Txt"`
}

func (e *Extension1) SetPlaceAndName(value string) {
	e.PlaceAndName = (*Max350Text)(&value)
}

func (e *Extension1) SetText(value string) {
	e.Text = (*Max350Text)(&value)
}
