package model

// Information about a discrepancy of a demand.
type Discrepancy1 struct {

	// Identification of the discrepancy.
	Identification *Max35Text `xml:"Id"`

	// Description of the discrepancy.
	Narrative *Max20000Text `xml:"Nrrtv"`
}

func (d *Discrepancy1) SetIdentification(value string) {
	d.Identification = (*Max35Text)(&value)
}

func (d *Discrepancy1) SetNarrative(value string) {
	d.Narrative = (*Max20000Text)(&value)
}
