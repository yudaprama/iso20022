package model

// Identification of the message number and the query identification.
type DocumentIdentification30 struct {

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber5Choice `xml:"MsgNb,omitempty"`

	// Reference to the query identification.
	Reference *Identification14 `xml:"Ref"`
}

func (d *DocumentIdentification30) AddMessageNumber() *DocumentNumber5Choice {
	d.MessageNumber = new(DocumentNumber5Choice)
	return d.MessageNumber
}

func (d *DocumentIdentification30) AddReference() *Identification14 {
	d.Reference = new(Identification14)
	return d.Reference
}
