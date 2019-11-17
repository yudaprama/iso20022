package model

// Identification of a document.
type DocumentNumber1 struct {

	// Number used to identify a message or document.
	Number *DocumentNumber1Choice `xml:"Nb"`
}

func (d *DocumentNumber1) AddNumber() *DocumentNumber1Choice {
	d.Number = new(DocumentNumber1Choice)
	return d.Number
}
