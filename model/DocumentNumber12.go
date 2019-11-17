package model

// Identification of the status being requested.
type DocumentNumber12 struct {

	// Number used to identify a message or document.
	Number *DocumentNumber5Choice `xml:"Nb"`

	// References of transaction for which the status is requested.
	References []*Identification15 `xml:"Refs"`
}

func (d *DocumentNumber12) AddNumber() *DocumentNumber5Choice {
	d.Number = new(DocumentNumber5Choice)
	return d.Number
}

func (d *DocumentNumber12) AddReferences() *Identification15 {
	newValue := new(Identification15)
	d.References = append(d.References, newValue)
	return newValue
}
