package model

// Data to authenticate.
type EncapsulatedContent1 struct {

	// Type of data which have been authenticated.
	ContentType *ContentType1Code `xml:"CnttTp"`

	// Actual data to authenticate.
	Content *Max10000Binary `xml:"Cntt,omitempty"`
}

func (e *EncapsulatedContent1) SetContentType(value string) {
	e.ContentType = (*ContentType1Code)(&value)
}

func (e *EncapsulatedContent1) SetContent(value string) {
	e.Content = (*Max10000Binary)(&value)
}
