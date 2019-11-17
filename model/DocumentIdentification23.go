package model

// Identifies a document by a unique identification and a date of issue.
type DocumentIdentification23 struct {

	// Identifies the document.
	Identification *Max35Text `xml:"Id"`

	// Date of issuance of the document.
	DateOfIssue *ISODate `xml:"DtOfIsse,omitempty"`

	// Identification of buyer order line item.
	OrderLineIdentification *Max35Text `xml:"OrdrLineId,omitempty"`
}

func (d *DocumentIdentification23) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *DocumentIdentification23) SetDateOfIssue(value string) {
	d.DateOfIssue = (*ISODate)(&value)
}

func (d *DocumentIdentification23) SetOrderLineIdentification(value string) {
	d.OrderLineIdentification = (*Max35Text)(&value)
}
