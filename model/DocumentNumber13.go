package model

// Identification of a document.
type DocumentNumber13 struct {

	// Number used to identify a message or document.
	Number *DocumentNumber5Choice `xml:"Nb"`
}

func (d *DocumentNumber13) AddNumber() *DocumentNumber5Choice {
	d.Number = new(DocumentNumber5Choice)
	return d.Number
}
