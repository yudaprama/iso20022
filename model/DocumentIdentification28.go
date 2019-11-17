package model

// Identifies a document by a unique identification and a date of issue.
type DocumentIdentification28 struct {

	// Identifies the document.
	Identification *Max35Text `xml:"Id,omitempty"`

	// Date of issuance of the document.
	DateOfIssue *ISODate `xml:"DtOfIsse"`
}

func (d *DocumentIdentification28) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *DocumentIdentification28) SetDateOfIssue(value string) {
	d.DateOfIssue = (*ISODate)(&value)
}
