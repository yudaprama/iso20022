package model

// Data to authenticate.
type EncapsulatedContent2 struct {

	// Type of data which have been authenticated.
	ContentType *ContentType1Code `xml:"CnttTp"`

	// Actual data to authenticate.
	Content *Max100KBinary `xml:"Cntt,omitempty"`
}

func (e *EncapsulatedContent2) SetContentType(value string) {
	e.ContentType = (*ContentType1Code)(&value)
}

func (e *EncapsulatedContent2) SetContent(value string) {
	e.Content = (*Max100KBinary)(&value)
}
