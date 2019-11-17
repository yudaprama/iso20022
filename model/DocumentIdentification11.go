package model

// Identification and creation date of a document.
type DocumentIdentification11 struct {

	// Unique identifier of the document (message) assigned by the sender of the document.
	Identification *Max35Text `xml:"Id"`

	// Date and time at which the document (message) was created by the sender.
	CreationDateTime *DateAndDateTimeChoice `xml:"CreDtTm,omitempty"`

	// Specifies if this document is a copy, a duplicate, or a duplicate of a copy.
	CopyDuplicate *CopyDuplicate1Code `xml:"CpyDplct,omitempty"`
}

func (d *DocumentIdentification11) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *DocumentIdentification11) AddCreationDateTime() *DateAndDateTimeChoice {
	d.CreationDateTime = new(DateAndDateTimeChoice)
	return d.CreationDateTime
}

func (d *DocumentIdentification11) SetCopyDuplicate(value string) {
	d.CopyDuplicate = (*CopyDuplicate1Code)(&value)
}
