package model

// Identifies a document by a unique identification and a date of issue.
type DocumentIdentification29 struct {

	// Identifies the document.
	Identification *Max35Text `xml:"Id"`

	// Date of issuance of the document.
	DateOfIssue *ISODate `xml:"DtOfIsse"`
}

func (d *DocumentIdentification29) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *DocumentIdentification29) SetDateOfIssue(value string) {
	d.DateOfIssue = (*ISODate)(&value)
}
