package model

// Identification of the message number and the query identification.
type DocumentIdentification48 struct {

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber6Choice `xml:"MsgNb,omitempty"`

	// Reference to the query identification.
	Reference *Identification16 `xml:"Ref"`
}

func (d *DocumentIdentification48) AddMessageNumber() *DocumentNumber6Choice {
	d.MessageNumber = new(DocumentNumber6Choice)
	return d.MessageNumber
}

func (d *DocumentIdentification48) AddReference() *Identification16 {
	d.Reference = new(Identification16)
	return d.Reference
}
