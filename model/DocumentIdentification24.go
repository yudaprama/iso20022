package model

// Identification of the message number and the query identification.
type DocumentIdentification24 struct {

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber1Choice `xml:"MsgNb,omitempty"`

	// Reference to the query identification.
	Reference *Identification1 `xml:"Ref"`
}

func (d *DocumentIdentification24) AddMessageNumber() *DocumentNumber1Choice {
	d.MessageNumber = new(DocumentNumber1Choice)
	return d.MessageNumber
}

func (d *DocumentIdentification24) AddReference() *Identification1 {
	d.Reference = new(Identification1)
	return d.Reference
}
