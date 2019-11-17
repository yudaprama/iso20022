package model

// Unique identifier of a document, message or transaction.
type Identification14 struct {

	// Unique identifier of a document, message or transaction.
	Identification *Max35Text `xml:"Id"`
}

func (i *Identification14) SetIdentification(value string) {
	i.Identification = (*Max35Text)(&value)
}
