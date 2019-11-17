package model

// Identification of the status being requested.
type DocumentNumber9 struct {

	// Number used to identify a message or document.
	Number *DocumentNumber1Choice `xml:"Nb"`

	// References of transaction for which the status is requested.
	References []*Identification11 `xml:"Refs"`
}

func (d *DocumentNumber9) AddNumber() *DocumentNumber1Choice {
	d.Number = new(DocumentNumber1Choice)
	return d.Number
}

func (d *DocumentNumber9) AddReferences() *Identification11 {
	newValue := new(Identification11)
	d.References = append(d.References, newValue)
	return newValue
}
