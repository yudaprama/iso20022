package model

// Identification of the status being requested.
type DocumentNumber15 struct {

	// Number used to identify a message or document.
	Number *DocumentNumber6Choice `xml:"Nb"`

	// References of transaction for which the status is requested.
	References []*Identification24 `xml:"Refs"`
}

func (d *DocumentNumber15) AddNumber() *DocumentNumber6Choice {
	d.Number = new(DocumentNumber6Choice)
	return d.Number
}

func (d *DocumentNumber15) AddReferences() *Identification24 {
	newValue := new(Identification24)
	d.References = append(d.References, newValue)
	return newValue
}
