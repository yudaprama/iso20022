package model

// Identification of a document.
type DocumentNumber14 struct {

	// Number used to identify a message or document.
	Number *DocumentNumber6Choice `xml:"Nb"`
}

func (d *DocumentNumber14) AddNumber() *DocumentNumber6Choice {
	d.Number = new(DocumentNumber6Choice)
	return d.Number
}
