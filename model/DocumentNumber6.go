package model

// Identification of the status being requested.
type DocumentNumber6 struct {

	// Number used to identify a message or document.
	Number *DocumentNumber1Choice `xml:"Nb"`

	// References of transaction for which the status is requested.
	References []*Identification7 `xml:"Refs"`
}

func (d *DocumentNumber6) AddNumber() *DocumentNumber1Choice {
	d.Number = new(DocumentNumber1Choice)
	return d.Number
}

func (d *DocumentNumber6) AddReferences() *Identification7 {
	newValue := new(Identification7)
	d.References = append(d.References, newValue)
	return newValue
}
