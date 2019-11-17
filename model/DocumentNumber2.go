package model

// Identification of the status being requested.
type DocumentNumber2 struct {

	// Number used to identify a message or document.
	Number *DocumentNumber1Choice `xml:"Nb"`

	// References of transaction for which the status is requested.
	References []*Identification2 `xml:"Refs"`
}

func (d *DocumentNumber2) AddNumber() *DocumentNumber1Choice {
	d.Number = new(DocumentNumber1Choice)
	return d.Number
}

func (d *DocumentNumber2) AddReferences() *Identification2 {
	newValue := new(Identification2)
	d.References = append(d.References, newValue)
	return newValue
}
