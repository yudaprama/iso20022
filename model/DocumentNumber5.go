package model

// Identification of the status being requested.
type DocumentNumber5 struct {

	// Number used to identify a message or document.
	Number *DocumentNumber1Choice `xml:"Nb"`

	// References of transaction for which the status is requested.
	References []*Identification6 `xml:"Refs"`
}

func (d *DocumentNumber5) AddNumber() *DocumentNumber1Choice {
	d.Number = new(DocumentNumber1Choice)
	return d.Number
}

func (d *DocumentNumber5) AddReferences() *Identification6 {
	newValue := new(Identification6)
	d.References = append(d.References, newValue)
	return newValue
}
