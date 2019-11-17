package model

// Generic identification scheme for a document.
type GenericDocumentIdentification5 struct {

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber6Choice `xml:"MsgNb,omitempty"`

	// Identification of the document.
	Identification *RestrictedFINXMax16Text `xml:"Id"`
}

func (g *GenericDocumentIdentification5) AddMessageNumber() *DocumentNumber6Choice {
	g.MessageNumber = new(DocumentNumber6Choice)
	return g.MessageNumber
}

func (g *GenericDocumentIdentification5) SetIdentification(value string) {
	g.Identification = (*RestrictedFINXMax16Text)(&value)
}
