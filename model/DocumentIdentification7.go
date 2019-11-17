package model

// Identifies a document by a unique identification and a date of issue.
type DocumentIdentification7 struct {

	// Identifies the document.
	Identification *Max35Text `xml:"Id"`

	// Date of issuance of the document.
	DateOfIssue *ISODate `xml:"DtOfIsse"`
}

func (d *DocumentIdentification7) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *DocumentIdentification7) SetDateOfIssue(value string) {
	d.DateOfIssue = (*ISODate)(&value)
}
