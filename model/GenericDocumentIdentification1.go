package model

// Generic identification scheme for a document.
type GenericDocumentIdentification1 struct {

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber1Choice `xml:"MsgNb,omitempty"`

	// Identification of the document.
	Identification *Max35Text `xml:"Id"`
}

func (g *GenericDocumentIdentification1) AddMessageNumber() *DocumentNumber1Choice {
	g.MessageNumber = new(DocumentNumber1Choice)
	return g.MessageNumber
}

func (g *GenericDocumentIdentification1) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}
