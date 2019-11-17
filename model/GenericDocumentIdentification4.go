package model

// Generic identification scheme for a document.
type GenericDocumentIdentification4 struct {

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber5Choice `xml:"MsgNb,omitempty"`

	// Identification of the document.
	Identification *Max35Text `xml:"Id"`
}

func (g *GenericDocumentIdentification4) AddMessageNumber() *DocumentNumber5Choice {
	g.MessageNumber = new(DocumentNumber5Choice)
	return g.MessageNumber
}

func (g *GenericDocumentIdentification4) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}
