package model

// Generic identification scheme for a document.
type GenericDocumentIdentification6 struct {

	// Message type number/message identifier of the message referenced in the linkage sequence.
	MessageNumber *DocumentNumber16Choice `xml:"MsgNb,omitempty"`

	// Identification of the document.
	Identification *RestrictedFINXMax16Text `xml:"Id"`
}

func (g *GenericDocumentIdentification6) AddMessageNumber() *DocumentNumber16Choice {
	g.MessageNumber = new(DocumentNumber16Choice)
	return g.MessageNumber
}

func (g *GenericDocumentIdentification6) SetIdentification(value string) {
	g.Identification = (*RestrictedFINXMax16Text)(&value)
}
