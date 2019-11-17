package model

// Identifies the document by providing a unique identification and optionally the date/time of the creation of the document.
type DocumentIdentification8 struct {

	// Unique identification of the document.
	Identification *Max35Text `xml:"Id"`

	// Date/time of the creation of the document.
	CreationDateTime *ISODateTime `xml:"CreDtTm,omitempty"`
}

func (d *DocumentIdentification8) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *DocumentIdentification8) SetCreationDateTime(value string) {
	d.CreationDateTime = (*ISODateTime)(&value)
}
