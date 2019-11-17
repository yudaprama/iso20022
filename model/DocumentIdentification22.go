package model

// Identifies a document by a unique identification and a date of issue.
type DocumentIdentification22 struct {

	// Identifies the document.
	Identification *Max35Text `xml:"Id"`

	// Date of issuance of the document.
	DateOfIssue *ISODate `xml:"DtOfIsse,omitempty"`
}

func (d *DocumentIdentification22) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *DocumentIdentification22) SetDateOfIssue(value string) {
	d.DateOfIssue = (*ISODate)(&value)
}
