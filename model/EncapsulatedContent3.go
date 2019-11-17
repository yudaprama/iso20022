package model

// Data to authenticate.
type EncapsulatedContent3 struct {

	// Type of data which have been authenticated.
	ContentType *ContentType2Code `xml:"CnttTp"`

	// Actual data to authenticate.
	Content *Max100KBinary `xml:"Cntt,omitempty"`
}

func (e *EncapsulatedContent3) SetContentType(value string) {
	e.ContentType = (*ContentType2Code)(&value)
}

func (e *EncapsulatedContent3) SetContent(value string) {
	e.Content = (*Max100KBinary)(&value)
}
